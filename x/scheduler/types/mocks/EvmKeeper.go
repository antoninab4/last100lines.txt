// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	types "github.com/palomachain/paloma/x/scheduler/types"
	mock "github.com/stretchr/testify/mock"

	xchain "github.com/palomachain/paloma/internal/x-chain"
)

// EvmKeeper is an autogenerated mock type for the EvmKeeper type
type EvmKeeper struct {
	mock.Mock
}

// PickValidatorForMessage provides a mock function with given fields: ctx, chainReferenceID, req
func (_m *EvmKeeper) PickValidatorForMessage(ctx context.Context, chainReferenceID string, req *xchain.JobRequirements) (string, error) {
	ret := _m.Called(ctx, chainReferenceID, req)

	if len(ret) == 0 {
		panic("no return value specified for PickValidatorForMessage")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *xchain.JobRequirements) (string, error)); ok {
		return rf(ctx, chainReferenceID, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *xchain.JobRequirements) string); ok {
		r0 = rf(ctx, chainReferenceID, req)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *xchain.JobRequirements) error); ok {
		r1 = rf(ctx, chainReferenceID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PreJobExecution provides a mock function with given fields: ctx, job
func (_m *EvmKeeper) PreJobExecution(ctx context.Context, job *types.Job) error {
	ret := _m.Called(ctx, job)

	if len(ret) == 0 {
		panic("no return value specified for PreJobExecution")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Job) error); ok {
		r0 = rf(ctx, job)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewEvmKeeper creates a new instance of EvmKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEvmKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *EvmKeeper {
	mock := &EvmKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}