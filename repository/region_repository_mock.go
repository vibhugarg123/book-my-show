// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository/region_repository.go

// Package repository is a generated GoMock package.
package repository

import (
	gomock "github.com/golang/mock/gomock"
	entities "github.com/vibhugarg123/book-my-show/entities"
	reflect "reflect"
)

// MockRegionRepository is a mock of RegionRepository interface
type MockRegionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRegionRepositoryMockRecorder
}

// MockRegionRepositoryMockRecorder is the mock recorder for MockRegionRepository
type MockRegionRepositoryMockRecorder struct {
	mock *MockRegionRepository
}

// NewMockRegionRepository creates a new mock instance
func NewMockRegionRepository(ctrl *gomock.Controller) *MockRegionRepository {
	mock := &MockRegionRepository{ctrl: ctrl}
	mock.recorder = &MockRegionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegionRepository) EXPECT() *MockRegionRepositoryMockRecorder {
	return m.recorder
}

// InsertRegion mocks base method
func (m *MockRegionRepository) InsertRegion(region entities.Region) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertRegion", region)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertRegion indicates an expected call of InsertRegion
func (mr *MockRegionRepositoryMockRecorder) InsertRegion(region interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertRegion", reflect.TypeOf((*MockRegionRepository)(nil).InsertRegion), region)
}

// FetchRegionById mocks base method
func (m *MockRegionRepository) FetchRegionById(arg0 int64) ([]entities.Region, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchRegionById", arg0)
	ret0, _ := ret[0].([]entities.Region)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchRegionById indicates an expected call of FetchRegionById
func (mr *MockRegionRepositoryMockRecorder) FetchRegionById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRegionById", reflect.TypeOf((*MockRegionRepository)(nil).FetchRegionById), arg0)
}