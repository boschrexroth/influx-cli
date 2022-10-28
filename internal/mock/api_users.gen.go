// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influxdata/influx-cli/v2/api (interfaces: UsersApi)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/influxdata/influx-cli/v2/api"
)

// MockUsersApi is a mock of UsersApi interface.
type MockUsersApi struct {
	ctrl     *gomock.Controller
	recorder *MockUsersApiMockRecorder
}

// MockUsersApiMockRecorder is the mock recorder for MockUsersApi.
type MockUsersApiMockRecorder struct {
	mock *MockUsersApi
}

// NewMockUsersApi creates a new mock instance.
func NewMockUsersApi(ctrl *gomock.Controller) *MockUsersApi {
	mock := &MockUsersApi{ctrl: ctrl}
	mock.recorder = &MockUsersApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersApi) EXPECT() *MockUsersApiMockRecorder {
	return m.recorder
}

// DeleteUsersID mocks base method.
func (m *MockUsersApi) DeleteUsersID(arg0 context.Context, arg1 string) api.ApiDeleteUsersIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUsersID", arg0, arg1)
	ret0, _ := ret[0].(api.ApiDeleteUsersIDRequest)
	return ret0
}

// DeleteUsersID indicates an expected call of DeleteUsersID.
func (mr *MockUsersApiMockRecorder) DeleteUsersID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUsersID", reflect.TypeOf((*MockUsersApi)(nil).DeleteUsersID), arg0, arg1)
}

// DeleteUsersIDExecute mocks base method.
func (m *MockUsersApi) DeleteUsersIDExecute(arg0 api.ApiDeleteUsersIDRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUsersIDExecute", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUsersIDExecute indicates an expected call of DeleteUsersIDExecute.
func (mr *MockUsersApiMockRecorder) DeleteUsersIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUsersIDExecute", reflect.TypeOf((*MockUsersApi)(nil).DeleteUsersIDExecute), arg0)
}

// DeleteUsersIDExecuteWithHttpInfo mocks base method.
func (m *MockUsersApi) DeleteUsersIDExecuteWithHttpInfo(arg0 api.ApiDeleteUsersIDRequest) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUsersIDExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUsersIDExecuteWithHttpInfo indicates an expected call of DeleteUsersIDExecuteWithHttpInfo.
func (mr *MockUsersApiMockRecorder) DeleteUsersIDExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUsersIDExecuteWithHttpInfo", reflect.TypeOf((*MockUsersApi)(nil).DeleteUsersIDExecuteWithHttpInfo), arg0)
}

// GetUsers mocks base method.
func (m *MockUsersApi) GetUsers(arg0 context.Context) api.ApiGetUsersRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", arg0)
	ret0, _ := ret[0].(api.ApiGetUsersRequest)
	return ret0
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUsersApiMockRecorder) GetUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUsersApi)(nil).GetUsers), arg0)
}

// GetUsersExecute mocks base method.
func (m *MockUsersApi) GetUsersExecute(arg0 api.ApiGetUsersRequest) (api.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersExecute", arg0)
	ret0, _ := ret[0].(api.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersExecute indicates an expected call of GetUsersExecute.
func (mr *MockUsersApiMockRecorder) GetUsersExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersExecute", reflect.TypeOf((*MockUsersApi)(nil).GetUsersExecute), arg0)
}

// GetUsersExecuteWithHttpInfo mocks base method.
func (m *MockUsersApi) GetUsersExecuteWithHttpInfo(arg0 api.ApiGetUsersRequest) (api.Users, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(api.Users)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUsersExecuteWithHttpInfo indicates an expected call of GetUsersExecuteWithHttpInfo.
func (mr *MockUsersApiMockRecorder) GetUsersExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersExecuteWithHttpInfo", reflect.TypeOf((*MockUsersApi)(nil).GetUsersExecuteWithHttpInfo), arg0)
}

// GetUsersID mocks base method.
func (m *MockUsersApi) GetUsersID(arg0 context.Context, arg1 string) api.ApiGetUsersIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersID", arg0, arg1)
	ret0, _ := ret[0].(api.ApiGetUsersIDRequest)
	return ret0
}

// GetUsersID indicates an expected call of GetUsersID.
func (mr *MockUsersApiMockRecorder) GetUsersID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersID", reflect.TypeOf((*MockUsersApi)(nil).GetUsersID), arg0, arg1)
}

// GetUsersIDExecute mocks base method.
func (m *MockUsersApi) GetUsersIDExecute(arg0 api.ApiGetUsersIDRequest) (api.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersIDExecute", arg0)
	ret0, _ := ret[0].(api.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersIDExecute indicates an expected call of GetUsersIDExecute.
func (mr *MockUsersApiMockRecorder) GetUsersIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersIDExecute", reflect.TypeOf((*MockUsersApi)(nil).GetUsersIDExecute), arg0)
}

// GetUsersIDExecuteWithHttpInfo mocks base method.
func (m *MockUsersApi) GetUsersIDExecuteWithHttpInfo(arg0 api.ApiGetUsersIDRequest) (api.UserResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersIDExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(api.UserResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUsersIDExecuteWithHttpInfo indicates an expected call of GetUsersIDExecuteWithHttpInfo.
func (mr *MockUsersApiMockRecorder) GetUsersIDExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersIDExecuteWithHttpInfo", reflect.TypeOf((*MockUsersApi)(nil).GetUsersIDExecuteWithHttpInfo), arg0)
}

// PatchUsersID mocks base method.
func (m *MockUsersApi) PatchUsersID(arg0 context.Context, arg1 string) api.ApiPatchUsersIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchUsersID", arg0, arg1)
	ret0, _ := ret[0].(api.ApiPatchUsersIDRequest)
	return ret0
}

// PatchUsersID indicates an expected call of PatchUsersID.
func (mr *MockUsersApiMockRecorder) PatchUsersID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchUsersID", reflect.TypeOf((*MockUsersApi)(nil).PatchUsersID), arg0, arg1)
}

// PatchUsersIDExecute mocks base method.
func (m *MockUsersApi) PatchUsersIDExecute(arg0 api.ApiPatchUsersIDRequest) (api.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchUsersIDExecute", arg0)
	ret0, _ := ret[0].(api.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchUsersIDExecute indicates an expected call of PatchUsersIDExecute.
func (mr *MockUsersApiMockRecorder) PatchUsersIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchUsersIDExecute", reflect.TypeOf((*MockUsersApi)(nil).PatchUsersIDExecute), arg0)
}

// PatchUsersIDExecuteWithHttpInfo mocks base method.
func (m *MockUsersApi) PatchUsersIDExecuteWithHttpInfo(arg0 api.ApiPatchUsersIDRequest) (api.UserResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchUsersIDExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(api.UserResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PatchUsersIDExecuteWithHttpInfo indicates an expected call of PatchUsersIDExecuteWithHttpInfo.
func (mr *MockUsersApiMockRecorder) PatchUsersIDExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchUsersIDExecuteWithHttpInfo", reflect.TypeOf((*MockUsersApi)(nil).PatchUsersIDExecuteWithHttpInfo), arg0)
}

// PostUsers mocks base method.
func (m *MockUsersApi) PostUsers(arg0 context.Context) api.ApiPostUsersRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostUsers", arg0)
	ret0, _ := ret[0].(api.ApiPostUsersRequest)
	return ret0
}

// PostUsers indicates an expected call of PostUsers.
func (mr *MockUsersApiMockRecorder) PostUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostUsers", reflect.TypeOf((*MockUsersApi)(nil).PostUsers), arg0)
}

// PostUsersExecute mocks base method.
func (m *MockUsersApi) PostUsersExecute(arg0 api.ApiPostUsersRequest) (api.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostUsersExecute", arg0)
	ret0, _ := ret[0].(api.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostUsersExecute indicates an expected call of PostUsersExecute.
func (mr *MockUsersApiMockRecorder) PostUsersExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostUsersExecute", reflect.TypeOf((*MockUsersApi)(nil).PostUsersExecute), arg0)
}

// PostUsersExecuteWithHttpInfo mocks base method.
func (m *MockUsersApi) PostUsersExecuteWithHttpInfo(arg0 api.ApiPostUsersRequest) (api.UserResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostUsersExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(api.UserResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PostUsersExecuteWithHttpInfo indicates an expected call of PostUsersExecuteWithHttpInfo.
func (mr *MockUsersApiMockRecorder) PostUsersExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostUsersExecuteWithHttpInfo", reflect.TypeOf((*MockUsersApi)(nil).PostUsersExecuteWithHttpInfo), arg0)
}

// PostUsersIDPassword mocks base method.
func (m *MockUsersApi) PostUsersIDPassword(arg0 context.Context, arg1 string) api.ApiPostUsersIDPasswordRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostUsersIDPassword", arg0, arg1)
	ret0, _ := ret[0].(api.ApiPostUsersIDPasswordRequest)
	return ret0
}

// PostUsersIDPassword indicates an expected call of PostUsersIDPassword.
func (mr *MockUsersApiMockRecorder) PostUsersIDPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostUsersIDPassword", reflect.TypeOf((*MockUsersApi)(nil).PostUsersIDPassword), arg0, arg1)
}

// PostUsersIDPasswordExecute mocks base method.
func (m *MockUsersApi) PostUsersIDPasswordExecute(arg0 api.ApiPostUsersIDPasswordRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostUsersIDPasswordExecute", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostUsersIDPasswordExecute indicates an expected call of PostUsersIDPasswordExecute.
func (mr *MockUsersApiMockRecorder) PostUsersIDPasswordExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostUsersIDPasswordExecute", reflect.TypeOf((*MockUsersApi)(nil).PostUsersIDPasswordExecute), arg0)
}

// PostUsersIDPasswordExecuteWithHttpInfo mocks base method.
func (m *MockUsersApi) PostUsersIDPasswordExecuteWithHttpInfo(arg0 api.ApiPostUsersIDPasswordRequest) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostUsersIDPasswordExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostUsersIDPasswordExecuteWithHttpInfo indicates an expected call of PostUsersIDPasswordExecuteWithHttpInfo.
func (mr *MockUsersApiMockRecorder) PostUsersIDPasswordExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostUsersIDPasswordExecuteWithHttpInfo", reflect.TypeOf((*MockUsersApi)(nil).PostUsersIDPasswordExecuteWithHttpInfo), arg0)
}

// PutUsersIDPassword mocks base method.
func (m *MockUsersApi) PutUsersIDPassword(arg0 context.Context, arg1 string) api.ApiPutUsersIDPasswordRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutUsersIDPassword", arg0, arg1)
	ret0, _ := ret[0].(api.ApiPutUsersIDPasswordRequest)
	return ret0
}

// PutUsersIDPassword indicates an expected call of PutUsersIDPassword.
func (mr *MockUsersApiMockRecorder) PutUsersIDPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutUsersIDPassword", reflect.TypeOf((*MockUsersApi)(nil).PutUsersIDPassword), arg0, arg1)
}

// PutUsersIDPasswordExecute mocks base method.
func (m *MockUsersApi) PutUsersIDPasswordExecute(arg0 api.ApiPutUsersIDPasswordRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutUsersIDPasswordExecute", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutUsersIDPasswordExecute indicates an expected call of PutUsersIDPasswordExecute.
func (mr *MockUsersApiMockRecorder) PutUsersIDPasswordExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutUsersIDPasswordExecute", reflect.TypeOf((*MockUsersApi)(nil).PutUsersIDPasswordExecute), arg0)
}

// PutUsersIDPasswordExecuteWithHttpInfo mocks base method.
func (m *MockUsersApi) PutUsersIDPasswordExecuteWithHttpInfo(arg0 api.ApiPutUsersIDPasswordRequest) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutUsersIDPasswordExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutUsersIDPasswordExecuteWithHttpInfo indicates an expected call of PutUsersIDPasswordExecuteWithHttpInfo.
func (mr *MockUsersApiMockRecorder) PutUsersIDPasswordExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutUsersIDPasswordExecuteWithHttpInfo", reflect.TypeOf((*MockUsersApi)(nil).PutUsersIDPasswordExecuteWithHttpInfo), arg0)
}
