// Code generated by MockGen. DO NOT EDIT.
// Source: userAppService.go

// Package mock_application is a generated GoMock package.
package mock_application

import (
	bizdto "douyin/types/bizdto"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockUserAppService is a mock of UserAppService interface.
type MockUserAppService struct {
	ctrl     *gomock.Controller
	recorder *MockUserAppServiceMockRecorder
}

// MockUserAppServiceMockRecorder is the mock recorder for MockUserAppService.
type MockUserAppServiceMockRecorder struct {
	mock *MockUserAppService
}

// NewMockUserAppService creates a new mock instance.
func NewMockUserAppService(ctrl *gomock.Controller) *MockUserAppService {
	mock := &MockUserAppService{ctrl: ctrl}
	mock.recorder = &MockUserAppServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAppService) EXPECT() *MockUserAppServiceMockRecorder {
	return m.recorder
}

// GetUser mocks base method.
func (m *MockUserAppService) GetUser(c *gin.Context, appUserID, userID int64) (*bizdto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", c, appUserID, userID)
	ret0, _ := ret[0].(*bizdto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserAppServiceMockRecorder) GetUser(c, appUserID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserAppService)(nil).GetUser), c, appUserID, userID)
}
