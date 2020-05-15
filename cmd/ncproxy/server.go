package main

import (
	"github.com/Microsoft/hcsshim/pkg/octtrpc"
	"github.com/containerd/ttrpc"
	"google.golang.org/grpc"
)

type server struct {
	ttrpc *ttrpc.Server
	grpc  *grpc.Server
}

func newServer() (*server, error) {
	ttrpcServer, err := ttrpc.NewServer(ttrpc.WithUnaryServerInterceptor(octtrpc.ServerInterceptor()))
	if err != nil {
		return nil, err
	}
	return &server{
		grpc:  grpc.NewServer(),
		ttrpc: ttrpcServer,
	}, nil
}
