package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/blck-snwmn/grpc-sample/processor"

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

	return &pb.RegisteredMessage{}, nil
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
	s := grpc.NewServer()
	pb.RegisterProcessorServer(s, &sv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
