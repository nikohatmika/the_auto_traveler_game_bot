// Code generated by mockery 2.9.4. DO NOT EDIT.

package mocks

import (
	equipments "auto_traveler/bussiness/equipments"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Find provides a mock function with given fields: ctx, eqType
func (_m *Repository) Find(ctx context.Context, eqType string) ([]equipments.Domain, error) {
	ret := _m.Called(ctx, eqType)

	var r0 []equipments.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) []equipments.Domain); ok {
		r0 = rf(ctx, eqType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]equipments.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, eqType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
