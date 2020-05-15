package main

import (
	"context"
	"net"
	"strings"

	"github.com/Microsoft/go-winio"
	"github.com/Microsoft/hcsshim/cmd/ncproxy/ncproxygrpc"
	"github.com/Microsoft/hcsshim/internal/log"
	"github.com/Microsoft/hcsshim/internal/ncproxyttrpc"
	"github.com/Microsoft/hcsshim/pkg/octtrpc"
	"github.com/containerd/ttrpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server struct {
	ttrpc *ttrpc.Server
	grpc  *grpc.Server
}

func newServer(ctx context.Context) (*server, error) {
	ttrpcServer, err := ttrpc.NewServer(ttrpc.WithUnaryServerInterceptor(octtrpc.ServerInterceptor()))
	if err != nil {
		log.G(ctx).WithError(err).Error("failed to create ttrpc server")
		return nil, err
	}
	return &server{
		grpc:  grpc.NewServer(),
		ttrpc: ttrpcServer,
	}, nil
}

func (s *server) setup(ctx context.Context) (net.Listener, net.Listener, error) {
	ncproxygrpc.RegisterNetworkConfigProxyServer(s.grpc, &grpcService{})
	ncproxyttrpc.RegisterNetworkConfigProxyService(s.ttrpc, &ttrpcService{})

	ttrpcListener, err := winio.ListenPipe(conf.TTRPCAddr, nil)
	if err != nil {
		log.G(ctx).WithError(err).Errorf("failed to listen on %s", conf.TTRPCAddr)
		return nil, nil, err
	}

	grpcListener, err := net.Listen("tcp", conf.GRPCAddr)
	if err != nil {
		log.G(ctx).WithError(err).Errorf("failed to listen on %s", conf.GRPCAddr)
		return nil, nil, err
	}
	return ttrpcListener, grpcListener, nil
}

func (s *server) gracefulShutdown(ctx context.Context) error {
	s.grpc.GracefulStop()
	return s.ttrpc.Shutdown(ctx)
}

func trapClosedConnErr(err error) error {
	if err == nil || strings.Contains(err.Error(), "use of closed network connection") {
		return nil
	}
	return err
}

func (s *server) serve(ctx context.Context, ttrpcListener net.Listener, grpcListener net.Listener, serveErr chan error) {
	go func() {
		log.G(ctx).WithFields(logrus.Fields{
			"address": conf.TTRPCAddr,
		}).Info("Serving ncproxy TTRPC service")

		serveErr <- trapClosedConnErr(s.ttrpc.Serve(ctx, ttrpcListener))
	}()

	go func() {
		log.G(ctx).WithFields(logrus.Fields{
			"address": conf.GRPCAddr,
		}).Info("Serving ncproxy GRPC service")

		defer grpcListener.Close()
		serveErr <- trapClosedConnErr(s.grpc.Serve(grpcListener))
	}()
}
