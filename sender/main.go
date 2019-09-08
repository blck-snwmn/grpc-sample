package main

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes"

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

	stream, err := client.RegisterStreamProcess(ctx)
	if err != nil {
		log.Fatalf("failed to do RegisterStreamProcess: %v", err)
	}
	defer stream.CloseSend()
	for i := 0; i < 10; i++ {
		//1ずつ送信し、レスポンスを受け取る
		err := stream.Send(&pb.Process{})
		if err != nil {
			log.Fatalf("failed to send proccess by RegisterStreamProcess's stream: %v", err)
		}
		resp, err := stream.Recv()
		if err != nil {
			log.Fatalf("failed to recieve response from  RegisterStreamProcess's stream: %v", err)
		}

		log.Printf("recieved at : %s", ptypes.TimestampString(resp.RecievedAt))
	}
	log.Println("end")
}
