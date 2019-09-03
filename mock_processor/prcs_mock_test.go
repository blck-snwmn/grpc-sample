package mock_processor_test

import (
	context "context"
	"testing"
	"time"

	prcs "github.com/blck-snwmn/grpc-sample/mock_processor"
	pb "github.com/blck-snwmn/grpc-sample/processor"
	"github.com/golang/mock/gomock"
)

func TestRegisterProcess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := prcs.NewMockProcessorClient(ctrl)

	process := &pb.Process{}

	client.EXPECT().RegisterProcess(
		gomock.Any(),
		process,
	).Return(&pb.RegisteredMessage{}, nil)

	testRegisterProcess(t, client)
}

func testRegisterProcess(t *testing.T, client pb.ProcessorClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := client.RegisterProcess(ctx, &pb.Process{})
	if err != nil {
		t.Errorf("mocking failed")
	}
}

var nof1 = &pb.Notification{Message: "message1"}
var nof2 = &pb.Notification{Message: "message2"}

func TestRequestNotification(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	stream := prcs.NewMockProcessor_RequestNotificationClient(ctrl)
	stream.EXPECT().Recv().Return(nof1, nil)
	stream.EXPECT().Recv().Return(nof2, nil)

	client := prcs.NewMockProcessorClient(ctrl)

	notif := &pb.NotificationRequest{}

	client.EXPECT().RequestNotification(
		gomock.Any(),
		notif,
	).Return(stream, nil)

	testRequestNotification(t, client)
}

func testRequestNotification(t *testing.T, client pb.ProcessorClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.RequestNotification(ctx, &pb.NotificationRequest{})
	if err != nil {
		t.Errorf("failed err=%v", err)
	}
	for _, nof := range []*pb.Notification{nof1, nof2} {
		got, err := stream.Recv()
		if err != nil {
			t.Errorf("failed to Recv. err=%v", err)
		}
		if got.Message != nof.Message {
			t.Errorf("failed to unmatch Recv. actual %v. want %v", got.Message, nof.Message)
		}
	}
}
