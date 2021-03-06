// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository/show_repository.go

// Package repository is a generated GoMock package.
package repository

import (
	gomock "github.com/golang/mock/gomock"
	entities "github.com/vibhugarg123/book-my-show/entities"
	reflect "reflect"
)

// MockShowRepository is a mock of ShowRepository interface
type MockShowRepository struct {
	ctrl     *gomock.Controller
	recorder *MockShowRepositoryMockRecorder
}

// MockShowRepositoryMockRecorder is the mock recorder for MockShowRepository
type MockShowRepositoryMockRecorder struct {
	mock *MockShowRepository
}

// NewMockShowRepository creates a new mock instance
func NewMockShowRepository(ctrl *gomock.Controller) *MockShowRepository {
	mock := &MockShowRepository{ctrl: ctrl}
	mock.recorder = &MockShowRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockShowRepository) EXPECT() *MockShowRepositoryMockRecorder {
	return m.recorder
}

// InsertShow mocks base method
func (m *MockShowRepository) InsertShow(arg0 *entities.Show) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertShow", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertShow indicates an expected call of InsertShow
func (mr *MockShowRepositoryMockRecorder) InsertShow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertShow", reflect.TypeOf((*MockShowRepository)(nil).InsertShow), arg0)
}

// FetchShowByMovieIdHallIdShowDate mocks base method
func (m *MockShowRepository) FetchShowByMovieIdHallIdShowDate(arg0 entities.Show) ([]entities.Show, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchShowByMovieIdHallIdShowDate", arg0)
	ret0, _ := ret[0].([]entities.Show)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchShowByMovieIdHallIdShowDate indicates an expected call of FetchShowByMovieIdHallIdShowDate
func (mr *MockShowRepositoryMockRecorder) FetchShowByMovieIdHallIdShowDate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchShowByMovieIdHallIdShowDate", reflect.TypeOf((*MockShowRepository)(nil).FetchShowByMovieIdHallIdShowDate), arg0)
}

// FetchShowByShowId mocks base method
func (m *MockShowRepository) FetchShowByShowId(arg0 int) ([]entities.Show, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchShowByShowId", arg0)
	ret0, _ := ret[0].([]entities.Show)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchShowByShowId indicates an expected call of FetchShowByShowId
func (mr *MockShowRepositoryMockRecorder) FetchShowByShowId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchShowByShowId", reflect.TypeOf((*MockShowRepository)(nil).FetchShowByShowId), arg0)
}

// UpdateSeatsByShowId mocks base method
func (m *MockShowRepository) UpdateSeatsByShowId(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSeatsByShowId", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSeatsByShowId indicates an expected call of UpdateSeatsByShowId
func (mr *MockShowRepositoryMockRecorder) UpdateSeatsByShowId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSeatsByShowId", reflect.TypeOf((*MockShowRepository)(nil).UpdateSeatsByShowId), arg0, arg1)
}
