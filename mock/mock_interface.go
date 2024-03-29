// Code generated by MockGen. DO NOT EDIT.
// Source: ../method-and-interface/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPrinter is a mock of Printer interface
type MockPrinter struct {
	ctrl     *gomock.Controller
	recorder *MockPrinterMockRecorder
}

// MockPrinterMockRecorder is the mock recorder for MockPrinter
type MockPrinterMockRecorder struct {
	mock *MockPrinter
}

// NewMockPrinter creates a new mock instance
func NewMockPrinter(ctrl *gomock.Controller) *MockPrinter {
	mock := &MockPrinter{ctrl: ctrl}
	mock.recorder = &MockPrinterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPrinter) EXPECT() *MockPrinterMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockPrinter) Add(arg0, arg1 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockPrinterMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockPrinter)(nil).Add), arg0, arg1)
}

// PlusTwo mocks base method
func (m *MockPrinter) PlusTwo(arg0 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlusTwo", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// PlusTwo indicates an expected call of PlusTwo
func (mr *MockPrinterMockRecorder) PlusTwo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlusTwo", reflect.TypeOf((*MockPrinter)(nil).PlusTwo), arg0)
}
