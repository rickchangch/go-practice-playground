package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "go-practice-playground/gRPC/normal/proto"

	"google.golang.org/grpc"
)

// - begin of block
// the below codes should be represented at the module handler/controller

type demoService struct {
	pb.UnimplementedDemoServiceServer
}

// New a service instance
var DemoInstance = new(demoService)

// Implement the function which has been defined in the protobuf service file
func (h *demoService) Action(ctx context.Context, pbReq *pb.Request) (*pb.Response, error) {
	if pbReq == nil {
		return nil, fmt.Errorf("null input")
	}

	log.Printf("successful grpc calling")

	return &pb.Response{}, nil
}

// - end of block

const port = ":50010"

// module main.go
func main() {
	// Listen port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// New gRPC server
	s := grpc.NewServer()

	// Can register multiple service on the same server
	pb.RegisterDemoServiceServer(s, DemoInstance)

	log.Printf("grpc starting serving..., listening on %v", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
