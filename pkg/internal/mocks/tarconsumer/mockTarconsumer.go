// Code generated by MockGen. DO NOT EDIT.
// Source: tarconsumer.go

package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockTarConsumer is a mock of TarConsumer interface
type MockTarConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockTarConsumerMockRecorder
}

// MockTarConsumerMockRecorder is the mock recorder for MockTarConsumer
type MockTarConsumerMockRecorder struct {
	mock *MockTarConsumer
}

// NewMockTarConsumer creates a new mock instance
func NewMockTarConsumer(ctrl *gomock.Controller) *MockTarConsumer {
	mock := &MockTarConsumer{ctrl: ctrl}
	mock.recorder = &MockTarConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockTarConsumer) EXPECT() *MockTarConsumerMockRecorder {
	return _m.recorder
}

// SetCtx mocks base method
func (_m *MockTarConsumer) SetCtx(ctx context.Context) {
	_m.ctrl.Call(_m, "SetCtx", ctx)
}

// SetCtx indicates an expected call of SetCtx
func (_mr *MockTarConsumerMockRecorder) SetCtx(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetCtx", reflect.TypeOf((*MockTarConsumer)(nil).SetCtx), arg0)
}

// SetURL mocks base method
func (_m *MockTarConsumer) SetURL(utl string) {
	_m.ctrl.Call(_m, "SetURL", utl)
}

// SetURL indicates an expected call of SetURL
func (_mr *MockTarConsumerMockRecorder) SetURL(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetURL", reflect.TypeOf((*MockTarConsumer)(nil).SetURL), arg0)
}

// SetHTTPClient mocks base method
func (_m *MockTarConsumer) SetHTTPClient(httpClient *http.Client) {
	_m.ctrl.Call(_m, "SetHTTPClient", httpClient)
}

// SetHTTPClient indicates an expected call of SetHTTPClient
func (_mr *MockTarConsumerMockRecorder) SetHTTPClient(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetHTTPClient", reflect.TypeOf((*MockTarConsumer)(nil).SetHTTPClient), arg0)
}

// GetTar mocks base method
func (_m *MockTarConsumer) GetTar(ctx context.Context) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "GetTar", ctx)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTar indicates an expected call of GetTar
func (_mr *MockTarConsumerMockRecorder) GetTar(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetTar", reflect.TypeOf((*MockTarConsumer)(nil).GetTar), arg0)
}
