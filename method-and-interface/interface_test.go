package mai

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/lj-211/go-test-example/mock"
)

func Test_InterfaceFunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockObj := mock.NewMockPrinter(ctrl)
	gomock.InOrder(
		mockObj.EXPECT().Add(gomock.Any(), gomock.Any()).Return(2),
		mockObj.EXPECT().PlusTwo(gomock.Any()).Return(4),
		mockObj.EXPECT().Add(1, 1).Return(4),
	)

	assert.Equal(t, 2, mockObj.Add(1, 2))
	assert.Equal(t, 4, mockObj.PlusTwo(1))
	assert.NotEqual(t, 2, mockObj.Add(1, 1))
}
