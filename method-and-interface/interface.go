package mai

import (
//"fmt"
)

// go get github.com/golang/mock/gomock
// go install github.com/golang/mock/mockgen

type Printer interface {
	Add(int, int) int
	PlusTwo(int) int
}

type NormalPrinter struct{}

func (self *NormalPrinter) Add(x int, y int) int {
	return x + y
}

func (self *NormalPrinter) PlusTwo(x int) int {
	return x * 2
}
