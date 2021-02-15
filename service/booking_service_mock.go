// Code generated by MockGen. DO NOT EDIT.
// Source: ./service/booking_service.go

// Package service is a generated GoMock package.
package service

import (
	gomock "github.com/golang/mock/gomock"
	entities "github.com/vibhugarg123/book-my-show/entities"
	reflect "reflect"
)

// MockBookingService is a mock of BookingService interface
type MockBookingService struct {
	ctrl     *gomock.Controller
	recorder *MockBookingServiceMockRecorder
}

// MockBookingServiceMockRecorder is the mock recorder for MockBookingService
type MockBookingServiceMockRecorder struct {
	mock *MockBookingService
}

// NewMockBookingService creates a new mock instance
func NewMockBookingService(ctrl *gomock.Controller) *MockBookingService {
	mock := &MockBookingService{ctrl: ctrl}
	mock.recorder = &MockBookingServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBookingService) EXPECT() *MockBookingServiceMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockBookingService) Add(arg0 entities.Booking) (entities.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(entities.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockBookingServiceMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockBookingService)(nil).Add), arg0)
}

// GetBooking mocks base method
func (m *MockBookingService) GetBooking(arg0 int) ([]entities.Booking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooking", arg0)
	ret0, _ := ret[0].([]entities.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBooking indicates an expected call of GetBooking
func (mr *MockBookingServiceMockRecorder) GetBooking(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooking", reflect.TypeOf((*MockBookingService)(nil).GetBooking), arg0)
}