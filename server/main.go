package main

import (
	"context"
	"log"
	"net"

	pb "github.com/blck-snwmn/grpc-sample/processor"

	"google.golang.org/grpc"
)

const (
	port = ":50055"
)

type server struct {
	recieved chan string
}

func (sv *server) RequestNotification(notif *pb.NotificationRequest, stream pb.Processor_RequestNotificationServer) error {
	for s := range sv.recieved {
		log.Println("send")
		err := stream.Send(&pb.Notification{Message: s})
		if err != nil {
			return err
		}
	}
	return nil
}

func (sv *server) RegisterProcess(ctx context.Context, p *pb.Process) (*pb.RegisteredMessage, error) {

	log.Println("register")

	// rand.Seed(time.Now().UnixNano())
	// n := rand.Intn(10)
	// time.Sleep(time.Duration(n) * time.Second)
	// message := fmt.Sprintf("hello:%d", n)
	// log.Println("register:" + message)
	// sv.recieved <- message
	sv.recieved <- "message"

	return &pb.RegisteredMessage{}, nil
}

func main() {
	sv := server{recieved: make(chan string, 9)}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProcessorServer(s, &sv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
