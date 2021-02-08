// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository/theatre_repository.go

// Package repository is a generated GoMock package.
package repository

import (
	gomock "github.com/golang/mock/gomock"
	entities "github.com/vibhugarg123/book-my-show/entities"
	reflect "reflect"
)

// MockTheatreRepository is a mock of TheatreRepository interface
type MockTheatreRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTheatreRepositoryMockRecorder
}

// MockTheatreRepositoryMockRecorder is the mock recorder for MockTheatreRepository
type MockTheatreRepositoryMockRecorder struct {
	mock *MockTheatreRepository
}

// NewMockTheatreRepository creates a new mock instance
func NewMockTheatreRepository(ctrl *gomock.Controller) *MockTheatreRepository {
	mock := &MockTheatreRepository{ctrl: ctrl}
	mock.recorder = &MockTheatreRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTheatreRepository) EXPECT() *MockTheatreRepositoryMockRecorder {
	return m.recorder
}

// InsertTheatre mocks base method
func (m *MockTheatreRepository) InsertTheatre(arg0 entities.Theatre) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTheatre", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertTheatre indicates an expected call of InsertTheatre
func (mr *MockTheatreRepositoryMockRecorder) InsertTheatre(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTheatre", reflect.TypeOf((*MockTheatreRepository)(nil).InsertTheatre), arg0)
}

// FetchTheatreByName mocks base method
func (m *MockTheatreRepository) FetchTheatreByName(arg0 string) ([]entities.Theatre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchTheatreByName", arg0)
	ret0, _ := ret[0].([]entities.Theatre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchTheatreByName indicates an expected call of FetchTheatreByName
func (mr *MockTheatreRepositoryMockRecorder) FetchTheatreByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchTheatreByName", reflect.TypeOf((*MockTheatreRepository)(nil).FetchTheatreByName), arg0)
}

// FetchTheatreByRegionId mocks base method
func (m *MockTheatreRepository) FetchTheatreByRegionId(arg0 int) ([]entities.Theatre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchTheatreByRegionId", arg0)
	ret0, _ := ret[0].([]entities.Theatre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchTheatreByRegionId indicates an expected call of FetchTheatreByRegionId
func (mr *MockTheatreRepositoryMockRecorder) FetchTheatreByRegionId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchTheatreByRegionId", reflect.TypeOf((*MockTheatreRepository)(nil).FetchTheatreByRegionId), arg0)
}

// FetchTheatreByNameRegionIdAndAddress mocks base method
func (m *MockTheatreRepository) FetchTheatreByNameRegionIdAndAddress(arg0 entities.Theatre) ([]entities.Theatre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchTheatreByNameRegionIdAndAddress", arg0)
	ret0, _ := ret[0].([]entities.Theatre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchTheatreByNameRegionIdAndAddress indicates an expected call of FetchTheatreByNameRegionIdAndAddress
func (mr *MockTheatreRepositoryMockRecorder) FetchTheatreByNameRegionIdAndAddress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchTheatreByNameRegionIdAndAddress", reflect.TypeOf((*MockTheatreRepository)(nil).FetchTheatreByNameRegionIdAndAddress), arg0)
}
