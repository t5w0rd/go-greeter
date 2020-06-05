package main

import (
	"google.golang.org/grpc"
	"greeter-grpc/handler"
	"log"
	"net"

	greeter "greeter-grpc/proto/greeter"
)

func main() {
	lis, err := net.Listen("tcp", ":50000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Server [grpc] Listening on %s", lis.Addr().String())

	s := grpc.NewServer()
	greeter.RegisterGreeterServer(s, &handler.Greeter{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
