// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository/hall_repository.go

// Package repository is a generated GoMock package.
package repository

import (
	gomock "github.com/golang/mock/gomock"
	entities "github.com/vibhugarg123/book-my-show/entities"
	reflect "reflect"
)

// MockHallRepository is a mock of HallRepository interface
type MockHallRepository struct {
	ctrl     *gomock.Controller
	recorder *MockHallRepositoryMockRecorder
}

// MockHallRepositoryMockRecorder is the mock recorder for MockHallRepository
type MockHallRepositoryMockRecorder struct {
	mock *MockHallRepository
}

// NewMockHallRepository creates a new mock instance
func NewMockHallRepository(ctrl *gomock.Controller) *MockHallRepository {
	mock := &MockHallRepository{ctrl: ctrl}
	mock.recorder = &MockHallRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHallRepository) EXPECT() *MockHallRepositoryMockRecorder {
	return m.recorder
}

// InsertHall mocks base method
func (m *MockHallRepository) InsertHall(arg0 entities.Hall) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertHall", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertHall indicates an expected call of InsertHall
func (mr *MockHallRepositoryMockRecorder) InsertHall(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertHall", reflect.TypeOf((*MockHallRepository)(nil).InsertHall), arg0)
}

// FetchHallByName mocks base method
func (m *MockHallRepository) FetchHallByName(arg0 string) ([]entities.Hall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchHallByName", arg0)
	ret0, _ := ret[0].([]entities.Hall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchHallByName indicates an expected call of FetchHallByName
func (mr *MockHallRepositoryMockRecorder) FetchHallByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchHallByName", reflect.TypeOf((*MockHallRepository)(nil).FetchHallByName), arg0)
}

// FetchHallByTheatreId mocks base method
func (m *MockHallRepository) FetchHallByTheatreId(arg0 int) ([]entities.Hall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchHallByTheatreId", arg0)
	ret0, _ := ret[0].([]entities.Hall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchHallByTheatreId indicates an expected call of FetchHallByTheatreId
func (mr *MockHallRepositoryMockRecorder) FetchHallByTheatreId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchHallByTheatreId", reflect.TypeOf((*MockHallRepository)(nil).FetchHallByTheatreId), arg0)
}

// FetchHallByHallId mocks base method
func (m *MockHallRepository) FetchHallByHallId(arg0 int64) ([]entities.Hall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchHallByHallId", arg0)
	ret0, _ := ret[0].([]entities.Hall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchHallByHallId indicates an expected call of FetchHallByHallId
func (mr *MockHallRepositoryMockRecorder) FetchHallByHallId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchHallByHallId", reflect.TypeOf((*MockHallRepository)(nil).FetchHallByHallId), arg0)
}

// FetchHallByNameAndTheatreId mocks base method
func (m *MockHallRepository) FetchHallByNameAndTheatreId(arg0 entities.Hall) ([]entities.Hall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchHallByNameAndTheatreId", arg0)
	ret0, _ := ret[0].([]entities.Hall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchHallByNameAndTheatreId indicates an expected call of FetchHallByNameAndTheatreId
func (mr *MockHallRepositoryMockRecorder) FetchHallByNameAndTheatreId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchHallByNameAndTheatreId", reflect.TypeOf((*MockHallRepository)(nil).FetchHallByNameAndTheatreId), arg0)
}
