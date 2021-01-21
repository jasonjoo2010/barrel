// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	store "github.com/projecteru2/barrel/store"
	mock "github.com/stretchr/testify/mock"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, encoder
func (_m *Store) Delete(ctx context.Context, encoder store.Encoder) (bool, error) {
	ret := _m.Called(ctx, encoder)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, store.Encoder) bool); ok {
		r0 = rf(ctx, encoder)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, store.Encoder) error); ok {
		r1 = rf(ctx, encoder)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, decoder
func (_m *Store) Get(ctx context.Context, decoder store.Decoder) (bool, error) {
	ret := _m.Called(ctx, decoder)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, store.Decoder) bool); ok {
		r0 = rf(ctx, decoder)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, store.Decoder) error); ok {
		r1 = rf(ctx, decoder)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAndDelete provides a mock function with given fields: ctx, decoder
func (_m *Store) GetAndDelete(ctx context.Context, decoder store.Decoder) (bool, error) {
	ret := _m.Called(ctx, decoder)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, store.Decoder) bool); ok {
		r0 = rf(ctx, decoder)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, store.Decoder) error); ok {
		r1 = rf(ctx, decoder)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: ctx, encoder
func (_m *Store) Put(ctx context.Context, encoder store.Encoder) error {
	ret := _m.Called(ctx, encoder)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, store.Encoder) error); ok {
		r0 = rf(ctx, encoder)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutMulti provides a mock function with given fields: ctx, encoders
func (_m *Store) PutMulti(ctx context.Context, encoders ...store.Encoder) error {
	_va := make([]interface{}, len(encoders))
	for _i := range encoders {
		_va[_i] = encoders[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...store.Encoder) error); ok {
		r0 = rf(ctx, encoders...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, encoder
func (_m *Store) Update(ctx context.Context, encoder store.Encoder) (bool, error) {
	ret := _m.Called(ctx, encoder)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, store.Encoder) bool); ok {
		r0 = rf(ctx, encoder)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, store.Encoder) error); ok {
		r1 = rf(ctx, encoder)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}