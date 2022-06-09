// Code generated by MockGen. DO NOT EDIT.
// Source: storage/local.go

// Package mock is a generated GoMock package.
package mock

import (
	entity "moviesapi/entity"
	external "moviesapi/external"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIStorage is a mock of IStorage interface.
type MockIStorage struct {
	ctrl     *gomock.Controller
	recorder *MockIStorageMockRecorder
}

// MockIStorageMockRecorder is the mock recorder for MockIStorage.
type MockIStorageMockRecorder struct {
	mock *MockIStorage
}

// NewMockIStorage creates a new mock instance.
func NewMockIStorage(ctrl *gomock.Controller) *MockIStorage {
	mock := &MockIStorage{ctrl: ctrl}
	mock.recorder = &MockIStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStorage) EXPECT() *MockIStorageMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIStorage) Create(data *external.MovieRequest) *entity.Movie {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", data)
	ret0, _ := ret[0].(*entity.Movie)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIStorageMockRecorder) Create(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIStorage)(nil).Create), data)
}

// DeleteByID mocks base method.
func (m *MockIStorage) DeleteByID(ID int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockIStorageMockRecorder) DeleteByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockIStorage)(nil).DeleteByID), ID)
}

// GetByID mocks base method.
func (m *MockIStorage) GetByID(ID int) *entity.Movie {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ID)
	ret0, _ := ret[0].(*entity.Movie)
	return ret0
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIStorageMockRecorder) GetByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIStorage)(nil).GetByID), ID)
}

// List mocks base method.
func (m *MockIStorage) List() []entity.Movie {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]entity.Movie)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockIStorageMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIStorage)(nil).List))
}

// UpdateByID mocks base method.
func (m *MockIStorage) UpdateByID(ID int, data *external.MovieRequest) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", ID, data)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UpdateByID indicates an expected call of UpdateByID.
func (mr *MockIStorageMockRecorder) UpdateByID(ID, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockIStorage)(nil).UpdateByID), ID, data)
}