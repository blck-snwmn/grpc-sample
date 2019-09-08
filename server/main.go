package main

import (
	"context"
	"io"
	"log"
	"net"
	"sync"

	pb "github.com/blck-snwmn/grpc-sample/processor"
	"github.com/golang/protobuf/ptypes"

	"google.golang.org/grpc"
)

const (
	port = ":50055"
)

type server struct {
	recieved      chan string
	mu            sync.Mutex
	notifications map[chan string]bool
}

func (sv *server) AddChan(ch chan string) {
	sv.mu.Lock()
	defer sv.mu.Unlock()

	sv.notifications[ch] = true
}

func (sv *server) RemoveChan(ch chan string) {
	sv.mu.Lock()
	defer sv.mu.Unlock()

	delete(sv.notifications, ch)
	close(ch)
}

func (sv *server) Notify() {
	for {
		select {
		case s := <-sv.recieved:
			func() {
				sv.mu.Lock()
				defer sv.mu.Unlock()
				for ch := range sv.notifications {
					ch <- s
				}
			}()
		}
	}
}

func (sv *server) RequestNotification(notif *pb.NotificationRequest, stream pb.Processor_RequestNotificationServer) error {
	ch := make(chan string)
	sv.AddChan(ch)
	defer sv.RemoveChan(ch)

	for s := range ch {
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

	sv.recieved <- "message"
	return &pb.RegisteredMessage{
		RecievedAt: ptypes.TimestampNow(),
	}, nil
}

func (sv *server) RegisterStreamProcess(stream pb.Processor_RegisterStreamProcessServer) error {
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			log.Println("RegisterStreamProcess recieve EOF")
			//とりあえず
			break
		}
		if err != nil {
			return err
		}

		// とりあえず気にせず投げる
		sv.recieved <- "message"
		err = stream.Send(&pb.RegisteredMessage{
			RecievedAt: ptypes.TimestampNow(),
		})
		if err != nil {
			return err
		}
	}
	log.Println("RegisterStreamProcess end")
	return nil
}

func main() {
	sv := server{
		recieved:      make(chan string, 9),
		notifications: make(map[chan string]bool),
	}
	go sv.Notify()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(doNothingIntercepter()),
		grpc.StreamInterceptor(doNothingStreamIntercepter()),
	)
	pb.RegisterProcessorServer(s, &sv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
