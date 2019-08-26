package mai

import (
	"fmt"
)

func GetNum(a, b int) int {
	fmt.Println("call GetNum")
	return a + b
}

func MethodTest(a, b int) int {
	return GetNum(a, b) + 100
}

type TestObj struct {
}

func (self *TestObj) LessZero(in int) bool {
	if in < 0 {
		return true
	}

	return false
}

func ObjTest() bool {
	obj := &TestObj{}
	return obj.LessZero(100)
}
