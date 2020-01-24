// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import graphql "github.com/machinebox/graphql"
import mock "github.com/stretchr/testify/mock"
import pkggraphql "github.com/kyma-incubator/compass/components/director/pkg/graphql"

// GraphQLRequestBuilder is an autogenerated mock type for the GraphQLRequestBuilder type
type GraphQLRequestBuilder struct {
	mock.Mock
}

// GetApplicationRequest provides a mock function with given fields: id
func (_m *GraphQLRequestBuilder) GetApplicationRequest(id string) *graphql.Request {
	ret := _m.Called(id)

	var r0 *graphql.Request
	if rf, ok := ret.Get(0).(func(string) *graphql.Request); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*graphql.Request)
		}
	}

	return r0
}

// RegisterApplicationRequest provides a mock function with given fields: input
func (_m *GraphQLRequestBuilder) RegisterApplicationRequest(input pkggraphql.ApplicationRegisterInput) (*graphql.Request, error) {
	ret := _m.Called(input)

	var r0 *graphql.Request
	if rf, ok := ret.Get(0).(func(pkggraphql.ApplicationRegisterInput) *graphql.Request); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*graphql.Request)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(pkggraphql.ApplicationRegisterInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnregisterApplicationRequest provides a mock function with given fields: id
func (_m *GraphQLRequestBuilder) UnregisterApplicationRequest(id string) *graphql.Request {
	ret := _m.Called(id)

	var r0 *graphql.Request
	if rf, ok := ret.Get(0).(func(string) *graphql.Request); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*graphql.Request)
		}
	}

	return r0
}
