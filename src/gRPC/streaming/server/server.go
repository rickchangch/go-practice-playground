package main

import (
	"context"
	"io"
	"log"
	"net"

	pb "go-practice-playground/gRPC/streaming/proto"

	"google.golang.org/grpc"
)

type StreamService struct {
	pb.UnimplementedStreamServiceServer
}

const (
	PORT = ":9002"
)

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterStreamServiceServer(server, &StreamService{})

	log.Printf("server listening at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *StreamService) GetFeature(ctx context.Context, req *pb.StreamRequest) (*pb.StreamResponse, error) {
	return &pb.StreamResponse{}, nil
}

// Server-side Streaming RPC
func (s *StreamService) ListFeatures(req *pb.StreamRequest, stream pb.StreamService_ListFeaturesServer) error {

	for n := 0; n <= 6; n++ {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  req.Pt.Name,
				Value: req.Pt.Value,
			},
		})

		if err != nil {
			return err
		}
	}

	return nil
}

// Client-side Streaming RPC
func (s *StreamService) RecordRoute(stream pb.StreamService_RecordRouteServer) error {
	for {
		r, err := stream.Recv()

		// 發現 stream close 後，回傳訊息，並通知 Client 也關閉 Stream
		if err == io.EOF {
			return stream.SendAndClose(&pb.StreamResponse{Pt: &pb.StreamPoint{Name: "gRPC Stream Server: Record", Value: 1}})
		}

		// error
		if err != nil {
			return err
		}

		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}
}

func (s *StreamService) RouteChat(stream pb.StreamService_RouteChatServer) error {
	return nil
}
