// Code generated by MockGen. DO NOT EDIT.
// Source: ./utils/time.go

// Package sgc7utils is a generated GoMock package.
package sanservutils

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockITime is a mock of ITime interface
type MockITime struct {
	ctrl     *gomock.Controller
	recorder *MockITimeMockRecorder
}

// MockITimeMockRecorder is the mock recorder for MockITime
type MockITimeMockRecorder struct {
	mock *MockITime
}

// NewMockITime creates a new mock instance
func NewMockITime(ctrl *gomock.Controller) *MockITime {
	mock := &MockITime{ctrl: ctrl}
	mock.recorder = &MockITimeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockITime) EXPECT() *MockITimeMockRecorder {
	return m.recorder
}

// Now mocks base method
func (m *MockITime) Now() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Now")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Now indicates an expected call of Now
func (mr *MockITimeMockRecorder) Now() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Now", reflect.TypeOf((*MockITime)(nil).Now))
}
