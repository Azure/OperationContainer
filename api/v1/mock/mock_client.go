// Code generated by MockGen. DO NOT EDIT.
// Source: dev.azure.com/service-hub-flg/service_hub/_git/service_hub.git/testing/canonical-output/operationcontainer/api/v1 (interfaces: OperationContainerClient)
//
// Generated by this command:
//
//	mockgen -package mock dev.azure.com/service-hub-flg/service_hub/_git/service_hub.git/testing/canonical-output/operationcontainer/api/v1 OperationContainerClient
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	v1 "github.com/Azure/OperationContainer/api/v1"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockOperationContainerClient is a mock of OperationContainerClient interface.
type MockOperationContainerClient struct {
	ctrl     *gomock.Controller
	recorder *MockOperationContainerClientMockRecorder
}

// MockOperationContainerClientMockRecorder is the mock recorder for MockOperationContainerClient.
type MockOperationContainerClientMockRecorder struct {
	mock *MockOperationContainerClient
}

// NewMockOperationContainerClient creates a new mock instance.
func NewMockOperationContainerClient(ctrl *gomock.Controller) *MockOperationContainerClient {
	mock := &MockOperationContainerClient{ctrl: ctrl}
	mock.recorder = &MockOperationContainerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperationContainerClient) EXPECT() *MockOperationContainerClientMockRecorder {
	return m.recorder
}

// CreateOperationStatus mocks base method.
func (m *MockOperationContainerClient) CreateOperationStatus(arg0 context.Context, arg1 *v1.CreateOperationStatusRequest, arg2 ...grpc.CallOption) (*v1.CreateOperationStatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOperationStatus", varargs...)
	ret0, _ := ret[0].(*v1.CreateOperationStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOperationStatus indicates an expected call of CreateOperationStatus.
func (mr *MockOperationContainerClientMockRecorder) CreateOperationStatus(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOperationStatus", reflect.TypeOf((*MockOperationContainerClient)(nil).CreateOperationStatus), varargs...)
}

// GetOperationEntity mocks base method.
func (m *MockOperationContainerClient) GetOperationEntity(arg0 context.Context, arg1 *v1.GetOperationEntityRequest, arg2 ...grpc.CallOption) (*v1.GetOperationEntityResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOperationEntity", varargs...)
	ret0, _ := ret[0].(*v1.GetOperationEntityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperationEntity indicates an expected call of GetOperationEntity.
func (mr *MockOperationContainerClientMockRecorder) GetOperationEntity(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperationEntity", reflect.TypeOf((*MockOperationContainerClient)(nil).GetOperationEntity), varargs...)
}

// GetOperationStatus mocks base method.
func (m *MockOperationContainerClient) GetOperationStatus(arg0 context.Context, arg1 *v1.GetOperationStatusRequest, arg2 ...grpc.CallOption) (*v1.GetOperationStatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOperationStatus", varargs...)
	ret0, _ := ret[0].(*v1.GetOperationStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperationStatus indicates an expected call of GetOperationStatus.
func (mr *MockOperationContainerClientMockRecorder) GetOperationStatus(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperationStatus", reflect.TypeOf((*MockOperationContainerClient)(nil).GetOperationStatus), varargs...)
}

// SayHello mocks base method.
func (m *MockOperationContainerClient) SayHello(arg0 context.Context, arg1 *v1.HelloRequest, arg2 ...grpc.CallOption) (*v1.HelloReply, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SayHello", varargs...)
	ret0, _ := ret[0].(*v1.HelloReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SayHello indicates an expected call of SayHello.
func (mr *MockOperationContainerClientMockRecorder) SayHello(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SayHello", reflect.TypeOf((*MockOperationContainerClient)(nil).SayHello), varargs...)
}

// UpdateOperationStatus mocks base method.
func (m *MockOperationContainerClient) UpdateOperationStatus(arg0 context.Context, arg1 *v1.UpdateOperationStatusRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateOperationStatus", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOperationStatus indicates an expected call of UpdateOperationStatus.
func (mr *MockOperationContainerClientMockRecorder) UpdateOperationStatus(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOperationStatus", reflect.TypeOf((*MockOperationContainerClient)(nil).UpdateOperationStatus), varargs...)
}