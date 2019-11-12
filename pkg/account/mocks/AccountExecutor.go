// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AccountExecutor is an autogenerated mock type for the AccountExecutor type
type AccountExecutor struct {
	mock.Mock
}

// AssumeRole provides a mock function with given fields: roleArn
func (_m *AccountExecutor) AssumeRole(roleArn string) error {
	ret := _m.Called(roleArn)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(roleArn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateRole provides a mock function with given fields:
func (_m *AccountExecutor) CreateRole() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
