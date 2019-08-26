package mai

import (
	"log"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
)

// NOTE
//		go test -run=Test_MethodTest -gcflags=-l
//		跑测试用例时需要关闭内联

// mock func
func MockGetNum(a, b int) int {
	log.Println("call MockGetNum")
	return 2
}

func Test_MethodTest(t *testing.T) {
	patches := gomonkey.ApplyFunc(GetNum, MockGetNum)
	defer patches.Reset()

	ret := MethodTest(1, 3)

	assert.Equal(t, 102, ret, "返回值必须是102")
}

// mock method
func Test_ObjTest(t *testing.T) {
	obj := &TestObj{}
	patches := gomonkey.ApplyMethod(reflect.TypeOf(obj), "LessZero", func(self *TestObj, in int) bool {
		return false
	})
	defer patches.Reset()
	assert.Equal(t, false, ObjTest(), "必须返回false")
}
