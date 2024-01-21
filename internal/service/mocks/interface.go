// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/interface.go

// Package mock_service is a generated GoMock package.
package mocks

import (
	context "context"
	models "github.com/bogdanpashtet/plata-currency-rates/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFrankfurterPrv is a mock of FrankfurterPrv interface.
type MockFrankfurterPrv struct {
	ctrl     *gomock.Controller
	recorder *MockFrankfurterPrvMockRecorder
}

// MockFrankfurterPrvMockRecorder is the mock recorder for MockFrankfurterPrv.
type MockFrankfurterPrvMockRecorder struct {
	mock *MockFrankfurterPrv
}

// NewMockFrankfurterPrv creates a new mock instance.
func NewMockFrankfurterPrv(ctrl *gomock.Controller) *MockFrankfurterPrv {
	mock := &MockFrankfurterPrv{ctrl: ctrl}
	mock.recorder = &MockFrankfurterPrvMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFrankfurterPrv) EXPECT() *MockFrankfurterPrvMockRecorder {
	return m.recorder
}

// GetRate mocks base method.
func (m *MockFrankfurterPrv) GetRate(ctx context.Context, toIso, fromIso string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRate", ctx, toIso, fromIso)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRate indicates an expected call of GetRate.
func (mr *MockFrankfurterPrvMockRecorder) GetRate(ctx, toIso, fromIso interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRate", reflect.TypeOf((*MockFrankfurterPrv)(nil).GetRate), ctx, toIso, fromIso)
}

// MockPostgres is a mock of Postgres interface.
type MockPostgres struct {
	ctrl     *gomock.Controller
	recorder *MockPostgresMockRecorder
}

// MockPostgresMockRecorder is the mock recorder for MockPostgres.
type MockPostgresMockRecorder struct {
	mock *MockPostgres
}

// NewMockPostgres creates a new mock instance.
func NewMockPostgres(ctrl *gomock.Controller) *MockPostgres {
	mock := &MockPostgres{ctrl: ctrl}
	mock.recorder = &MockPostgresMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostgres) EXPECT() *MockPostgresMockRecorder {
	return m.recorder
}

// AddToQueue mocks base method.
func (m *MockPostgres) AddToQueue(ctx context.Context, rate models.CurrencyRate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToQueue", ctx, rate)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToQueue indicates an expected call of AddToQueue.
func (mr *MockPostgresMockRecorder) AddToQueue(ctx, rate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToQueue", reflect.TypeOf((*MockPostgres)(nil).AddToQueue), ctx, rate)
}

// ConfirmQueue mocks base method.
func (m *MockPostgres) ConfirmQueue(ctx context.Context) (models.CurrencyRateWithDt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfirmQueue", ctx)
	ret0, _ := ret[0].(models.CurrencyRateWithDt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfirmQueue indicates an expected call of ConfirmQueue.
func (mr *MockPostgresMockRecorder) ConfirmQueue(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmQueue", reflect.TypeOf((*MockPostgres)(nil).ConfirmQueue), ctx)
}

// GetById mocks base method.
func (m *MockPostgres) GetById(ctx context.Context, id string) (models.CurrencyRateWithDt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(models.CurrencyRateWithDt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockPostgresMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockPostgres)(nil).GetById), ctx, id)
}

// GetLastRate mocks base method.
func (m *MockPostgres) GetLastRate(ctx context.Context, toIso, fromIso string) (models.CurrencyRateLast, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastRate", ctx, toIso, fromIso)
	ret0, _ := ret[0].(models.CurrencyRateLast)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastRate indicates an expected call of GetLastRate.
func (mr *MockPostgresMockRecorder) GetLastRate(ctx, toIso, fromIso interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastRate", reflect.TypeOf((*MockPostgres)(nil).GetLastRate), ctx, toIso, fromIso)
}