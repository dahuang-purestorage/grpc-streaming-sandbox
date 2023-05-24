package main

import (
	"log"
	"net"

	"github.com/dahuang-purestorage/grpc-streaming-sandbox/pkg/adder"
	"github.com/dahuang-purestorage/grpc-streaming-sandbox/pkg/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	// Create new gRPC server instance
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}

	// Register gRPC server
	api.RegisterAdderServer(s, srv)
	// reflection.Register(srv)

	// Listen on port 8080
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// Start gRPC server
	logrus.Infof("Start serving on port 8080")
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
