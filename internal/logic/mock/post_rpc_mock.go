// Code generated by MockGen. DO NOT EDIT.
// Source: post_rpc.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	postrpc "github.com/xh-polaris/meowchat-post-rpc/postrpc"
	grpc "google.golang.org/grpc"
)

// MockPostRpc is a mock of PostRpc interface.
type MockPostRpc struct {
	ctrl     *gomock.Controller
	recorder *MockPostRpcMockRecorder
}

// MockPostRpcMockRecorder is the mock recorder for MockPostRpc.
type MockPostRpcMockRecorder struct {
	mock *MockPostRpc
}

// NewMockPostRpc creates a new mock instance.
func NewMockPostRpc(ctrl *gomock.Controller) *MockPostRpc {
	mock := &MockPostRpc{ctrl: ctrl}
	mock.recorder = &MockPostRpcMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostRpc) EXPECT() *MockPostRpcMockRecorder {
	return m.recorder
}

// CreatePost mocks base method.
func (m *MockPostRpc) CreatePost(ctx context.Context, in *postrpc.CreatePostReq, opts ...grpc.CallOption) (*postrpc.CreatePostResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreatePost", varargs...)
	ret0, _ := ret[0].(*postrpc.CreatePostResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockPostRpcMockRecorder) CreatePost(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockPostRpc)(nil).CreatePost), varargs...)
}

// DeletePost mocks base method.
func (m *MockPostRpc) DeletePost(ctx context.Context, in *postrpc.DeletePostReq, opts ...grpc.CallOption) (*postrpc.DeletePostResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeletePost", varargs...)
	ret0, _ := ret[0].(*postrpc.DeletePostResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockPostRpcMockRecorder) DeletePost(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockPostRpc)(nil).DeletePost), varargs...)
}

// ListPost mocks base method.
func (m *MockPostRpc) ListPost(ctx context.Context, in *postrpc.ListPostReq, opts ...grpc.CallOption) (*postrpc.ListPostResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPost", varargs...)
	ret0, _ := ret[0].(*postrpc.ListPostResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPost indicates an expected call of ListPost.
func (mr *MockPostRpcMockRecorder) ListPost(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPost", reflect.TypeOf((*MockPostRpc)(nil).ListPost), varargs...)
}

// ListPostByUserId mocks base method.
func (m *MockPostRpc) ListPostByUserId(ctx context.Context, in *postrpc.ListPostByUserIdReq, opts ...grpc.CallOption) (*postrpc.ListPostByUserIdResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPostByUserId", varargs...)
	ret0, _ := ret[0].(*postrpc.ListPostByUserIdResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPostByUserId indicates an expected call of ListPostByUserId.
func (mr *MockPostRpcMockRecorder) ListPostByUserId(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPostByUserId", reflect.TypeOf((*MockPostRpc)(nil).ListPostByUserId), varargs...)
}

// RetrievePost mocks base method.
func (m *MockPostRpc) RetrievePost(ctx context.Context, in *postrpc.RetrievePostReq, opts ...grpc.CallOption) (*postrpc.RetrievePostResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrievePost", varargs...)
	ret0, _ := ret[0].(*postrpc.RetrievePostResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrievePost indicates an expected call of RetrievePost.
func (mr *MockPostRpcMockRecorder) RetrievePost(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrievePost", reflect.TypeOf((*MockPostRpc)(nil).RetrievePost), varargs...)
}

// SearchPost mocks base method.
func (m *MockPostRpc) SearchPost(ctx context.Context, in *postrpc.SearchPostReq, opts ...grpc.CallOption) (*postrpc.SearchPostResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchPost", varargs...)
	ret0, _ := ret[0].(*postrpc.SearchPostResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPost indicates an expected call of SearchPost.
func (mr *MockPostRpcMockRecorder) SearchPost(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPost", reflect.TypeOf((*MockPostRpc)(nil).SearchPost), varargs...)
}

// UpdatePost mocks base method.
func (m *MockPostRpc) UpdatePost(ctx context.Context, in *postrpc.UpdatePostReq, opts ...grpc.CallOption) (*postrpc.UpdatePostResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdatePost", varargs...)
	ret0, _ := ret[0].(*postrpc.UpdatePostResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockPostRpcMockRecorder) UpdatePost(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockPostRpc)(nil).UpdatePost), varargs...)
}
