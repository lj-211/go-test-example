package nettest

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/lj-211/go-test-example/mock"
	"github.com/lj-211/go-test-example/nettest/proto"
)

func Test_CallGrpcClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockObj := mock.NewMockHelloClient(ctrl)
	gomock.InOrder(
		mockObj.EXPECT().SayHello(context.Background(), gomock.Any()).Return(
			&proto.HelloRes{
				Id: 1234,
			}, nil,
		),
	)

	res, err := mockObj.SayHello(context.Background(), &proto.HelloReq{
		Msg: "测试",
	})
	assert.Nil(t, err)
	assert.Equal(t, uint32(1234), res.Id, "返回值必须是1234")
}
