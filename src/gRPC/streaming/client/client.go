package main

import (
	"context"
	"io"
	"log"

	pb "go-practice-playground/gRPC/streaming/proto"

	"google.golang.org/grpc"
)

const (
	PORT = ":9002"
)

func main() {
	conn, err := grpc.Dial(PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewStreamServiceClient(conn)

	err = printGet(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRPC Stream Client: Get", Value: 2018}})
	if err != nil {
		log.Fatalf("printLists.err: %v", err)
	}

	err = printLists(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRPC Stream Client: List", Value: 2018}})
	if err != nil {
		log.Fatalf("printLists.err: %v", err)
	}

	err = printRecord(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRPC Stream Client: Record", Value: 2018}})
	if err != nil {
		log.Fatalf("printRecord.err: %v", err)
	}

	err = printRoute(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRPC Stream Client: Route", Value: 2018}})
	if err != nil {
		log.Fatalf("printRoute.err: %v", err)
	}
}

func printGet(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	return nil
}

// Server-side Streaming RPC
func printLists(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.ListFeatures(context.Background(), r)
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()

		// indicates successful end of stream.
		if err == io.EOF {
			break
		}

		// fail err
		if err != nil {
			return err
		}

		log.Printf("resp: pt.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)

	}
	return nil
}

// Client-side Streaming RPC
func printRecord(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.RecordRoute(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n <= 6; n++ {
		err := stream.Send(r)
		if err != nil {
			return err
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Printf("resp: pt.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)

	return nil
}

func printRoute(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	return nil
}
