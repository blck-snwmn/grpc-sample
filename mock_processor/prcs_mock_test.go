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
