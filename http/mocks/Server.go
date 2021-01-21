// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	http "github.com/projecteru2/barrel/http"
	mock "github.com/stretchr/testify/mock"
)

// Server is an autogenerated mock type for the Server type
type Server struct {
	mock.Mock
}

// Close provides a mock function with given fields: _a0
func (_m *Server) Close(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CloseAsync provides a mock function with given fields: _a0
func (_m *Server) CloseAsync(_a0 func(error)) {
	_m.Called(_a0)
}

// ServeHTTP provides a mock function with given fields: _a0
func (_m *Server) ServeHTTP(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServeHTTPS provides a mock function with given fields: _a0, _a1
func (_m *Server) ServeHTTPS(_a0 string, _a1 http.TLSConfig) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, http.TLSConfig) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServeUnix provides a mock function with given fields: _a0, _a1
func (_m *Server) ServeUnix(_a0 string, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}