package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"bitcoin-address-cycles"
	"bitcoin-address-cycles/rpc"
)

// Start the server.
func main() {
	solutionServer := bitcoin_address_cycles.NewSolutionServer()
	grpcServer := grpc.NewServer()
	rpc.RegisterSolutionServer(grpcServer, solutionServer)

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("cannot create listener: %v", err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("error running grpc server: %v", err)
	}
}
