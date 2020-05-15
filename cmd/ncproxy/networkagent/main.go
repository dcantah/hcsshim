package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Microsoft/go-winio/pkg/guid"

	"github.com/Microsoft/hcsshim/cmd/ncproxy/configagent"
	"github.com/Microsoft/hcsshim/cmd/ncproxy/ncproxygrpc"
	"github.com/Microsoft/hcsshim/internal/hcsoci"
	"github.com/Microsoft/hcsshim/internal/log"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// This is a barebones example of an implementation of the network
// config agent service that ncproxy talks to. This is solely used to test.
const (
	netName     = "ContainerPlat-nat"
	listenAddr  = "127.0.0.1:9201"
	ncProxyAddr = "127.0.0.1:9200"
)

type service struct {
	client ncproxygrpc.NetworkConfigProxyClient
}

func (s *service) ConnectNamespaceToNetwork(ctx context.Context, req *configagent.ConnectNamespaceToNetworkRequest) (*configagent.ConnectNamespaceToNetworkResponse, error) {
	// For testing we pass in the network name instead as it's easier for local use
	// and to look up.
	log.G(ctx).WithFields(logrus.Fields{
		"namespace": req.NamespaceID,
		"networkID": req.NetworkID,
	}).Info("ConnectNamespaceToNetwork request")
	endpoints, err := hcsoci.GetNamespaceEndpoints(ctx, req.NamespaceID)
	if err != nil {
		return nil, fmt.Errorf("failed to get namespace endpoints: %s", err)
	}
	for _, endpoint := range endpoints {
		if endpoint.VirtualNetworkName == netName {
			nicID, err := guid.NewV4()
			if err != nil {
				return nil, fmt.Errorf("failed to create nic GUID: %s", err)
			}
			nsReq := &ncproxygrpc.AddNICRequest{
				NamespaceID: req.NamespaceID,
				NicID:       nicID.String(),
				EndpointID:  endpoint.Id,
			}
			if _, err := s.client.AddNIC(ctx, nsReq); err != nil {
				return nil, err
			}
		}
	}
	return &configagent.ConnectNamespaceToNetworkResponse{}, nil
}

func main() {
	ctx := context.Background()

	sigChan := make(chan os.Signal, 1)
	serveErr := make(chan error, 1)
	defer close(serveErr)
	signal.Notify(sigChan, syscall.SIGINT)
	defer signal.Stop(sigChan)

	grpcClient, err := grpc.Dial(ncProxyAddr, grpc.WithInsecure())
	if err != nil {
		log.G(ctx).WithError(err).Error("failed to connect to ncproxy")
		os.Exit(1)
	}
	defer grpcClient.Close()

	log.G(ctx).WithField("addr", ncProxyAddr).Info("connected to ncproxy")
	ncproxyClient := ncproxygrpc.NewNetworkConfigProxyClient(grpcClient)
	service := &service{ncproxyClient}
	server := grpc.NewServer()
	configagent.RegisterNetworkConfigAgentServer(server, service)

	grpcListener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.G(ctx).WithError(err).Errorf("failed to listen on %s", grpcListener.Addr().String())
		os.Exit(1)
	}

	go func() {
		defer grpcListener.Close()
		if err := server.Serve(grpcListener); err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				serveErr <- nil
			}
			serveErr <- err
		}
	}()

	// Wait for server error or user cancellation.
	select {
	case <-sigChan:
		log.G(ctx).Info("Received interrupt. Closing")
		break
	case err := <-serveErr:
		if err != nil {
			log.G(ctx).WithError(err).Fatal("grpc service failure")
		}
	}

	// Cancel inflight requests and shutdown service
	server.GracefulStop()
}
