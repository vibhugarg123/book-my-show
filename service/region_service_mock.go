// Code generated by MockGen. DO NOT EDIT.
// Source: ./service/region_service.go

// Package service is a generated GoMock package.
package service

import (
	gomock "github.com/golang/mock/gomock"
	entities "github.com/vibhugarg123/book-my-show/entities"
	reflect "reflect"
)

// MockRegionService is a mock of RegionService interface
type MockRegionService struct {
	ctrl     *gomock.Controller
	recorder *MockRegionServiceMockRecorder
}

// MockRegionServiceMockRecorder is the mock recorder for MockRegionService
type MockRegionServiceMockRecorder struct {
	mock *MockRegionService
}

// NewMockRegionService creates a new mock instance
func NewMockRegionService(ctrl *gomock.Controller) *MockRegionService {
	mock := &MockRegionService{ctrl: ctrl}
	mock.recorder = &MockRegionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegionService) EXPECT() *MockRegionServiceMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockRegionService) Add(arg0 entities.Region) (entities.Region, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(entities.Region)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockRegionServiceMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockRegionService)(nil).Add), arg0)
}

// GetRegionById mocks base method
func (m *MockRegionService) GetRegionById(arg0 int) (entities.Region, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegionById", arg0)
	ret0, _ := ret[0].(entities.Region)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegionById indicates an expected call of GetRegionById
func (mr *MockRegionServiceMockRecorder) GetRegionById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegionById", reflect.TypeOf((*MockRegionService)(nil).GetRegionById), arg0)
}
