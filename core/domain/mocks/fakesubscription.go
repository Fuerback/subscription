// Code generated by MockGen. DO NOT EDIT.
// Source: ./core/domain/subscription.go

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	domain "github.com/Fuerback/subscription/core/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockSubscriptionService is a mock of SubscriptionService interface.
type MockSubscriptionService struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionServiceMockRecorder
}

// MockSubscriptionServiceMockRecorder is the mock recorder for MockSubscriptionService.
type MockSubscriptionServiceMockRecorder struct {
	mock *MockSubscriptionService
}

// NewMockSubscriptionService creates a new mock instance.
func NewMockSubscriptionService(ctrl *gomock.Controller) *MockSubscriptionService {
	mock := &MockSubscriptionService{ctrl: ctrl}
	mock.recorder = &MockSubscriptionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscriptionService) EXPECT() *MockSubscriptionServiceMockRecorder {
	return m.recorder
}

// FetchOne mocks base method.
func (m *MockSubscriptionService) FetchOne(response http.ResponseWriter, request *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FetchOne", response, request)
}

// FetchOne indicates an expected call of FetchOne.
func (mr *MockSubscriptionServiceMockRecorder) FetchOne(response, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchOne", reflect.TypeOf((*MockSubscriptionService)(nil).FetchOne), response, request)
}

// MockSubscriptionUseCase is a mock of SubscriptionUseCase interface.
type MockSubscriptionUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionUseCaseMockRecorder
}

// MockSubscriptionUseCaseMockRecorder is the mock recorder for MockSubscriptionUseCase.
type MockSubscriptionUseCaseMockRecorder struct {
	mock *MockSubscriptionUseCase
}

// NewMockSubscriptionUseCase creates a new mock instance.
func NewMockSubscriptionUseCase(ctrl *gomock.Controller) *MockSubscriptionUseCase {
	mock := &MockSubscriptionUseCase{ctrl: ctrl}
	mock.recorder = &MockSubscriptionUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscriptionUseCase) EXPECT() *MockSubscriptionUseCaseMockRecorder {
	return m.recorder
}

// FetchOne mocks base method.
func (m *MockSubscriptionUseCase) FetchOne(id string) (*domain.SubscriptionDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchOne", id)
	ret0, _ := ret[0].(*domain.SubscriptionDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchOne indicates an expected call of FetchOne.
func (mr *MockSubscriptionUseCaseMockRecorder) FetchOne(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchOne", reflect.TypeOf((*MockSubscriptionUseCase)(nil).FetchOne), id)
}

// MockSubscriptionRepository is a mock of SubscriptionRepository interface.
type MockSubscriptionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionRepositoryMockRecorder
}

// MockSubscriptionRepositoryMockRecorder is the mock recorder for MockSubscriptionRepository.
type MockSubscriptionRepositoryMockRecorder struct {
	mock *MockSubscriptionRepository
}

// NewMockSubscriptionRepository creates a new mock instance.
func NewMockSubscriptionRepository(ctrl *gomock.Controller) *MockSubscriptionRepository {
	mock := &MockSubscriptionRepository{ctrl: ctrl}
	mock.recorder = &MockSubscriptionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscriptionRepository) EXPECT() *MockSubscriptionRepositoryMockRecorder {
	return m.recorder
}

// FetchOne mocks base method.
func (m *MockSubscriptionRepository) FetchOne(id string) (*domain.SubscriptionDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchOne", id)
	ret0, _ := ret[0].(*domain.SubscriptionDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchOne indicates an expected call of FetchOne.
func (mr *MockSubscriptionRepositoryMockRecorder) FetchOne(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchOne", reflect.TypeOf((*MockSubscriptionRepository)(nil).FetchOne), id)
}
