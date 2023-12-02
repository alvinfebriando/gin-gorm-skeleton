// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	valueobject "github.com/alvinfebriando/gin-gorm-skeleton/valueobject"
)

// BaseRepository is an autogenerated mock type for the BaseRepository type
type BaseRepository[T interface{}] struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, t
func (_m *BaseRepository[T]) Create(ctx context.Context, t *T) (*T, error) {
	ret := _m.Called(ctx, t)

	var r0 *T
	if rf, ok := ret.Get(0).(func(context.Context, *T) *T); ok {
		r0 = rf(ctx, t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *T) error); ok {
		r1 = rf(ctx, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, t
func (_m *BaseRepository[T]) Delete(ctx context.Context, t *T) error {
	ret := _m.Called(ctx, t)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *T) error); ok {
		r0 = rf(ctx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctx, query
func (_m *BaseRepository[T]) Find(ctx context.Context, query *valueobject.Query) ([]*T, error) {
	ret := _m.Called(ctx, query)

	var r0 []*T
	if rf, ok := ret.Get(0).(func(context.Context, *valueobject.Query) []*T); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*T)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *valueobject.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: ctx, id
func (_m *BaseRepository[T]) FindById(ctx context.Context, id uint) (*T, error) {
	ret := _m.Called(ctx, id)

	var r0 *T
	if rf, ok := ret.Get(0).(func(context.Context, uint) *T); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: ctx, query
func (_m *BaseRepository[T]) FindOne(ctx context.Context, query *valueobject.Query) (*T, error) {
	ret := _m.Called(ctx, query)

	var r0 *T
	if rf, ok := ret.Get(0).(func(context.Context, *valueobject.Query) *T); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *valueobject.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, t
func (_m *BaseRepository[T]) Update(ctx context.Context, t *T) (*T, error) {
	ret := _m.Called(ctx, t)

	var r0 *T
	if rf, ok := ret.Get(0).(func(context.Context, *T) *T); ok {
		r0 = rf(ctx, t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *T) error); ok {
		r1 = rf(ctx, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBaseRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewBaseRepository creates a new instance of BaseRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBaseRepository[T interface{}](t mockConstructorTestingTNewBaseRepository) *BaseRepository[T] {
	mock := &BaseRepository[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
