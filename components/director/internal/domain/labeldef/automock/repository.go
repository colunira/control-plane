// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import context "context"

import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-incubator/compass/components/director/internal/model"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, def
func (_m *Repository) Create(ctx context.Context, def model.LabelDefinition) error {
	ret := _m.Called(ctx, def)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.LabelDefinition) error); ok {
		r0 = rf(ctx, def)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}