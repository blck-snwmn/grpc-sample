package main

import (
	"context"
	"io"
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

	client := pb.NewProcessorClient(conn)

	// use RequestNotification
	// server streaming RPC
	stream, err := client.RequestNotification(context.Background(), &pb.NotificationRequest{})
	if err != nil {
		log.Fatalf("RequestNotification failed: %v", err)
	}
	for {
		notif, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.RequestNotification(_) = _, %v", client, err)
		}
		log.Printf("recieved: %v\n", notif)
	}
	log.Println("end")
}
