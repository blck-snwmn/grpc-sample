package main

import (
	"context"
	"log"

	pb "github.com/blck-snwmn/grpc-sample/processor"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50055"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()

	client := pb.NewProcessorClient(conn)

	client.RegisterProcess(ctx, &pb.Process{})
	log.Println("end")
}
