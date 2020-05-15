package main

import (
	"context"
	"flag"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Microsoft/go-winio"
	"github.com/Microsoft/hcsshim/cmd/ncproxy/ncproxygrpc"
	"github.com/Microsoft/hcsshim/internal/ncproxyttrpc"

	"github.com/Microsoft/hcsshim/cmd/ncproxy/configagent"
	"github.com/Microsoft/hcsshim/internal/computeagent"
	"github.com/Microsoft/hcsshim/internal/log"
	"github.com/sirupsen/logrus"
)

var (
	configPath = flag.String("config", "", "Path to JSON configuration file.")
	// Global mapping of network namespace ID to shim compute agent ttrpc service.
	namespaceToShim = make(map[string]computeagent.ComputeAgentService)
	// Mapping of network name to config agent clients
	networkToConfigAgent = make(map[string]configagent.NetworkConfigAgentClient)
)

func main() {
	flag.Parse()
	ctx := context.Background()
	config, err := loadConfig(*configPath)
	if err != nil {
		log.G(ctx).WithError(err).Error("failed getting configuration file")
		os.Exit(1)
	}

	if config.GRPCAddr == "" {
		log.G(ctx).Error("missing GRPC endpoint in config")
		os.Exit(1)
	}
	if config.TTRPCAddr == "" {
		log.G(ctx).Error("missing TTRPC endpoint in config")
		os.Exit(1)
	}

	// Construct config agent client connections from config file information.
	if err := configToClients(ctx, config); err != nil {
		log.G(ctx).WithError(err).Error("failed to dial clients")
		os.Exit(1)
	}

	log.G(ctx).Info("Starting ncproxy")

	sigChan := make(chan os.Signal, 1)
	serveErr := make(chan error, 1)
	defer close(serveErr)
	signal.Notify(sigChan, syscall.SIGINT)
	defer signal.Stop(sigChan)

	// Create new server and then register NetworkConfigProxyServices.
	server, err := newServer()
	if err != nil {
		log.G(ctx).WithError(err).Error("failed to create new server")
	}

	ncproxygrpc.RegisterNetworkConfigProxyServer(server.grpc, &grpcService{})
	ncproxyttrpc.RegisterNetworkConfigProxyService(server.ttrpc, &ttrpcService{})

	ttrpcListener, err := winio.ListenPipe(config.TTRPCAddr, nil)
	if err != nil {
		log.G(ctx).WithError(err).Errorf("failed to listen on %s", ttrpcListener.Addr().String())
		os.Exit(1)
	}

	grpcListener, err := net.Listen("tcp", config.GRPCAddr)
	if err != nil {
		log.G(ctx).WithError(err).Errorf("failed to listen on %s", grpcListener.Addr().String())
		os.Exit(1)
	}

	go func() {
		log.G(ctx).WithFields(logrus.Fields{
			"address": config.TTRPCAddr,
		}).Info("Serving TTRPC service")

		// ttrpc.Serve already closes the listener when returning
		if err := server.ttrpc.Serve(ctx, ttrpcListener); err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				serveErr <- nil
			}
			serveErr <- err
		}
	}()

	go func() {
		log.G(ctx).WithFields(logrus.Fields{
			"address": config.GRPCAddr,
		}).Info("Serving GRPC service")

		defer grpcListener.Close()
		if err := server.grpc.Serve(grpcListener); err != nil {
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
	case err := <-serveErr:
		if err != nil {
			log.G(ctx).WithError(err).Fatal("service failure")
		}
	}

	// Cancel inflight requests and shutdown services
	server.ttrpc.Shutdown(ctx)
	server.grpc.GracefulStop()
}
