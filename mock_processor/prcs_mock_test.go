package mock_processor_test

import (
	context "context"
	"testing"
	"time"

	prcs "github.com/blck-snwmn/grpc-sample/mock_processor"
	pb "github.com/blck-snwmn/grpc-sample/processor"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"
)

var expectedRM = &pb.RegisteredMessage{}

func TestRegisterProcess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := prcs.NewMockProcessorClient(ctrl)

	process := &pb.Process{}

	time, _ := ptypes.TimestampProto(time.Now())
	expectedRM.RecievedAt = time

	client.EXPECT().RegisterProcess(
		gomock.Any(),
		process,
	).Return(expectedRM, nil)

	testRegisterProcess(t, client)
}

func testRegisterProcess(t *testing.T, client pb.ProcessorClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.RegisterProcess(ctx, &pb.Process{})
	if err != nil {
		t.Errorf("mocking failed")
	}
	if resp.RecievedAt != expectedRM.RecievedAt {
		t.Errorf("RecievedAt doesn't equal response.RecievedAt")
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

func TestRegisterStreamProcess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	stream := prcs.NewMockProcessor_RegisterStreamProcessClient(ctrl)
	stream.EXPECT().Send(gomock.Any()).Return(nil)
	stream.EXPECT().Recv().Return(&pb.RegisteredMessage{}, nil)

	client := prcs.NewMockProcessorClient(ctrl)
	client.EXPECT().RegisterStreamProcess(
		gomock.Any(),
	).Return(stream, nil)
	testRegisterStreamProcess(t, client)
}

func testRegisterStreamProcess(t *testing.T, client pb.ProcessorClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.RegisterStreamProcess(ctx)
	if err != nil {
		t.Errorf("failed err=%v", err)
	}
	if err = stream.Send(&pb.Process{}); err != nil {
		t.Errorf("failed to Send. err=%v", err)
	}
	_, err = stream.Recv()
	if err != nil {
		t.Errorf("failed to Recv. err=%v", err)
	}
}
