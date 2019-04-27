// Code generated by MockGen. DO NOT EDIT.
// Source: app/repo/book_repo.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	repo "github.com/typical-go/typical-rest-server/app/repo"
	reflect "reflect"
)

// MockBookRepository is a mock of BookRepository interface
type MockBookRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBookRepositoryMockRecorder
}

// MockBookRepositoryMockRecorder is the mock recorder for MockBookRepository
type MockBookRepositoryMockRecorder struct {
	mock *MockBookRepository
}

// NewMockBookRepository creates a new mock instance
func NewMockBookRepository(ctrl *gomock.Controller) *MockBookRepository {
	mock := &MockBookRepository{ctrl: ctrl}
	mock.recorder = &MockBookRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBookRepository) EXPECT() *MockBookRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockBookRepository) Get(id int64) (*repo.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*repo.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockBookRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBookRepository)(nil).Get), id)
}

// List mocks base method
func (m *MockBookRepository) List() ([]*repo.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*repo.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockBookRepositoryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockBookRepository)(nil).List))
}

// Insert mocks base method
func (m *MockBookRepository) Insert(book repo.Book) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", book)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockBookRepositoryMockRecorder) Insert(book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockBookRepository)(nil).Insert), book)
}

// Delete mocks base method
func (m *MockBookRepository) Delete(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockBookRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBookRepository)(nil).Delete), id)
}

// Update mocks base method
func (m *MockBookRepository) Update(book repo.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", book)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockBookRepositoryMockRecorder) Update(book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBookRepository)(nil).Update), book)
}