// Code generated by MockGen. DO NOT EDIT.
// Source: ./service/movie_service.go

// Package service is a generated GoMock package.
package service

import (
	gomock "github.com/golang/mock/gomock"
	entities "github.com/vibhugarg123/book-my-show/entities"
	reflect "reflect"
)

// MockMovieService is a mock of MovieService interface
type MockMovieService struct {
	ctrl     *gomock.Controller
	recorder *MockMovieServiceMockRecorder
}

// MockMovieServiceMockRecorder is the mock recorder for MockMovieService
type MockMovieServiceMockRecorder struct {
	mock *MockMovieService
}

// NewMockMovieService creates a new mock instance
func NewMockMovieService(ctrl *gomock.Controller) *MockMovieService {
	mock := &MockMovieService{ctrl: ctrl}
	mock.recorder = &MockMovieServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMovieService) EXPECT() *MockMovieServiceMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockMovieService) Add(arg0 entities.Movie) (entities.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(entities.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockMovieServiceMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockMovieService)(nil).Add), arg0)
}

// GetActiveMovies mocks base method
func (m *MockMovieService) GetActiveMovies() ([]entities.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveMovies")
	ret0, _ := ret[0].([]entities.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveMovies indicates an expected call of GetActiveMovies
func (mr *MockMovieServiceMockRecorder) GetActiveMovies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveMovies", reflect.TypeOf((*MockMovieService)(nil).GetActiveMovies))
}
