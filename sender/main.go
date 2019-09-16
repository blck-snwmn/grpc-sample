package main

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"

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

	// use go-grpc-middleware/auth
	// auth middleware's token
	header := metadata.Pairs("authorization", "bearer p@ssword")
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	client := pb.NewProcessorClient(conn)

	// use RegisterProcess
	// unary RPC for client
	_, err = client.RegisterProcess(ctx, &pb.Process{})
	if err != nil {
		log.Fatalf("failed to do RegisterProcess: %v", err)
	}

	// use RegisterStreamProcess
	// bidirectional streaming RPC
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
