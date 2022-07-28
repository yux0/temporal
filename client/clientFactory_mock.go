// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: clientfactory.go

// Package client is a generated GoMock package.
package client

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	v1 "go.temporal.io/api/workflowservice/v1"
	v10 "go.temporal.io/server/api/adminservice/v1"
	v11 "go.temporal.io/server/api/historyservice/v1"
	v12 "go.temporal.io/server/api/matchingservice/v1"
	common "go.temporal.io/server/common"
	dynamicconfig "go.temporal.io/server/common/dynamicconfig"
	log "go.temporal.io/server/common/log"
	membership "go.temporal.io/server/common/membership"
	metrics "go.temporal.io/server/common/metrics"
)

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// NewAdminClientWithTimeout mocks base method.
func (m *MockFactory) NewAdminClientWithTimeout(rpcAddress string, timeout, largeTimeout time.Duration) v10.AdminServiceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAdminClientWithTimeout", rpcAddress, timeout, largeTimeout)
	ret0, _ := ret[0].(v10.AdminServiceClient)
	return ret0
}

// NewAdminClientWithTimeout indicates an expected call of NewAdminClientWithTimeout.
func (mr *MockFactoryMockRecorder) NewAdminClientWithTimeout(rpcAddress, timeout, largeTimeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAdminClientWithTimeout", reflect.TypeOf((*MockFactory)(nil).NewAdminClientWithTimeout), rpcAddress, timeout, largeTimeout)
}

// NewFrontendClient mocks base method.
func (m *MockFactory) NewFrontendClient(rpcAddress string) (v1.WorkflowServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewFrontendClient", rpcAddress)
	ret0, _ := ret[0].(v1.WorkflowServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewFrontendClient indicates an expected call of NewFrontendClient.
func (mr *MockFactoryMockRecorder) NewFrontendClient(rpcAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewFrontendClient", reflect.TypeOf((*MockFactory)(nil).NewFrontendClient), rpcAddress)
}

// NewFrontendClientWithTimeout mocks base method.
func (m *MockFactory) NewFrontendClientWithTimeout(rpcAddress string, timeout, longPollTimeout time.Duration) v1.WorkflowServiceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewFrontendClientWithTimeout", rpcAddress, timeout, longPollTimeout)
	ret0, _ := ret[0].(v1.WorkflowServiceClient)
	return ret0
}

// NewFrontendClientWithTimeout indicates an expected call of NewFrontendClientWithTimeout.
func (mr *MockFactoryMockRecorder) NewFrontendClientWithTimeout(rpcAddress, timeout, longPollTimeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewFrontendClientWithTimeout", reflect.TypeOf((*MockFactory)(nil).NewFrontendClientWithTimeout), rpcAddress, timeout, longPollTimeout)
}

// NewHistoryClient mocks base method.
func (m *MockFactory) NewHistoryClient() (v11.HistoryServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewHistoryClient")
	ret0, _ := ret[0].(v11.HistoryServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewHistoryClient indicates an expected call of NewHistoryClient.
func (mr *MockFactoryMockRecorder) NewHistoryClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewHistoryClient", reflect.TypeOf((*MockFactory)(nil).NewHistoryClient))
}

// NewHistoryClientWithTimeout mocks base method.
func (m *MockFactory) NewHistoryClientWithTimeout(timeout time.Duration) (v11.HistoryServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewHistoryClientWithTimeout", timeout)
	ret0, _ := ret[0].(v11.HistoryServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewHistoryClientWithTimeout indicates an expected call of NewHistoryClientWithTimeout.
func (mr *MockFactoryMockRecorder) NewHistoryClientWithTimeout(timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewHistoryClientWithTimeout", reflect.TypeOf((*MockFactory)(nil).NewHistoryClientWithTimeout), timeout)
}

// NewLocalAdminClientWithTimeout mocks base method.
func (m *MockFactory) NewLocalAdminClientWithTimeout(timeout, largeTimeout time.Duration) (v10.AdminServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewLocalAdminClientWithTimeout", timeout, largeTimeout)
	ret0, _ := ret[0].(v10.AdminServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewLocalAdminClientWithTimeout indicates an expected call of NewLocalAdminClientWithTimeout.
func (mr *MockFactoryMockRecorder) NewLocalAdminClientWithTimeout(timeout, largeTimeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewLocalAdminClientWithTimeout", reflect.TypeOf((*MockFactory)(nil).NewLocalAdminClientWithTimeout), timeout, largeTimeout)
}

// NewLocalFrontendClientWithTimeout mocks base method.
func (m *MockFactory) NewLocalFrontendClientWithTimeout(timeout, longPollTimeout time.Duration) (v1.WorkflowServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewLocalFrontendClientWithTimeout", timeout, longPollTimeout)
	ret0, _ := ret[0].(v1.WorkflowServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewLocalFrontendClientWithTimeout indicates an expected call of NewLocalFrontendClientWithTimeout.
func (mr *MockFactoryMockRecorder) NewLocalFrontendClientWithTimeout(timeout, longPollTimeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewLocalFrontendClientWithTimeout", reflect.TypeOf((*MockFactory)(nil).NewLocalFrontendClientWithTimeout), timeout, longPollTimeout)
}

// NewMatchingClient mocks base method.
func (m *MockFactory) NewMatchingClient(namespaceIDToName NamespaceIDToNameFunc) (v12.MatchingServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewMatchingClient", namespaceIDToName)
	ret0, _ := ret[0].(v12.MatchingServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewMatchingClient indicates an expected call of NewMatchingClient.
func (mr *MockFactoryMockRecorder) NewMatchingClient(namespaceIDToName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMatchingClient", reflect.TypeOf((*MockFactory)(nil).NewMatchingClient), namespaceIDToName)
}

// NewMatchingClientWithTimeout mocks base method.
func (m *MockFactory) NewMatchingClientWithTimeout(namespaceIDToName NamespaceIDToNameFunc, timeout, longPollTimeout time.Duration) (v12.MatchingServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewMatchingClientWithTimeout", namespaceIDToName, timeout, longPollTimeout)
	ret0, _ := ret[0].(v12.MatchingServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewMatchingClientWithTimeout indicates an expected call of NewMatchingClientWithTimeout.
func (mr *MockFactoryMockRecorder) NewMatchingClientWithTimeout(namespaceIDToName, timeout, longPollTimeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMatchingClientWithTimeout", reflect.TypeOf((*MockFactory)(nil).NewMatchingClientWithTimeout), namespaceIDToName, timeout, longPollTimeout)
}

// MockFactoryProvider is a mock of FactoryProvider interface.
type MockFactoryProvider struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryProviderMockRecorder
}

// MockFactoryProviderMockRecorder is the mock recorder for MockFactoryProvider.
type MockFactoryProviderMockRecorder struct {
	mock *MockFactoryProvider
}

// NewMockFactoryProvider creates a new mock instance.
func NewMockFactoryProvider(ctrl *gomock.Controller) *MockFactoryProvider {
	mock := &MockFactoryProvider{ctrl: ctrl}
	mock.recorder = &MockFactoryProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactoryProvider) EXPECT() *MockFactoryProviderMockRecorder {
	return m.recorder
}

// NewFactory mocks base method.
func (m *MockFactoryProvider) NewFactory(rpcFactory common.RPCFactory, monitor membership.Monitor, metricsClient metrics.Client, dc *dynamicconfig.Collection, numberOfHistoryShards int32, logger, throttledLogger log.Logger) Factory {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewFactory", rpcFactory, monitor, metricsClient, dc, numberOfHistoryShards, logger, throttledLogger)
	ret0, _ := ret[0].(Factory)
	return ret0
}

// NewFactory indicates an expected call of NewFactory.
func (mr *MockFactoryProviderMockRecorder) NewFactory(rpcFactory, monitor, metricsClient, dc, numberOfHistoryShards, logger, throttledLogger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewFactory", reflect.TypeOf((*MockFactoryProvider)(nil).NewFactory), rpcFactory, monitor, metricsClient, dc, numberOfHistoryShards, logger, throttledLogger)
}
