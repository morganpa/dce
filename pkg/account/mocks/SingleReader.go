// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/Optum/dce/pkg/model"

// SingleReader is an autogenerated mock type for the SingleReader type
type SingleReader struct {
	mock.Mock
}

// GetAccountByID provides a mock function with given fields: accountID, _a1
func (_m *SingleReader) GetAccountByID(accountID string, _a1 *model.Account) error {
	ret := _m.Called(accountID, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *model.Account) error); ok {
		r0 = rf(accountID, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}