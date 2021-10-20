package main

import (
	"context"
	"fmt"
	"log"

	pb "go-practice-playground/gRPC/normal/proto"

	"google.golang.org/grpc"
)

func main() {

	// gRPC 預設是要用加密的，但此處沒有加密的相關設定，因此用 Insecure
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("dial grpc server error: %v", err))
	}
	defer conn.Close()

	// 建立 service client
	client := pb.NewDemoServiceClient(conn)

	// 藉由 client 發出 grpc 請求
	resp, err := client.Action(context.TODO(), &pb.Request{Name: "Bob"})

	if err != nil {
		log.Println("grpc function error: Action")
	}

	log.Println(resp)
	log.Println("end...")
}
