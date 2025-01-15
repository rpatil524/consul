// Code generated by mockery v2.41.0. DO NOT EDIT.

package mockpbdataplane

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	pbdataplane "github.com/hashicorp/consul/proto-public/pbdataplane"
)

// DataplaneServiceClient is an autogenerated mock type for the DataplaneServiceClient type
type DataplaneServiceClient struct {
	mock.Mock
}

type DataplaneServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *DataplaneServiceClient) EXPECT() *DataplaneServiceClient_Expecter {
	return &DataplaneServiceClient_Expecter{mock: &_m.Mock}
}

// GetEnvoyBootstrapParams provides a mock function with given fields: ctx, in, opts
func (_m *DataplaneServiceClient) GetEnvoyBootstrapParams(ctx context.Context, in *pbdataplane.GetEnvoyBootstrapParamsRequest, opts ...grpc.CallOption) (*pbdataplane.GetEnvoyBootstrapParamsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetEnvoyBootstrapParams")
	}

	var r0 *pbdataplane.GetEnvoyBootstrapParamsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pbdataplane.GetEnvoyBootstrapParamsRequest, ...grpc.CallOption) (*pbdataplane.GetEnvoyBootstrapParamsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pbdataplane.GetEnvoyBootstrapParamsRequest, ...grpc.CallOption) *pbdataplane.GetEnvoyBootstrapParamsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pbdataplane.GetEnvoyBootstrapParamsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pbdataplane.GetEnvoyBootstrapParamsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DataplaneServiceClient_GetEnvoyBootstrapParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEnvoyBootstrapParams'
type DataplaneServiceClient_GetEnvoyBootstrapParams_Call struct {
	*mock.Call
}

// GetEnvoyBootstrapParams is a helper method to define mock.On call
//   - ctx context.Context
//   - in *pbdataplane.GetEnvoyBootstrapParamsRequest
//   - opts ...grpc.CallOption
func (_e *DataplaneServiceClient_Expecter) GetEnvoyBootstrapParams(ctx interface{}, in interface{}, opts ...interface{}) *DataplaneServiceClient_GetEnvoyBootstrapParams_Call {
	return &DataplaneServiceClient_GetEnvoyBootstrapParams_Call{Call: _e.mock.On("GetEnvoyBootstrapParams",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *DataplaneServiceClient_GetEnvoyBootstrapParams_Call) Run(run func(ctx context.Context, in *pbdataplane.GetEnvoyBootstrapParamsRequest, opts ...grpc.CallOption)) *DataplaneServiceClient_GetEnvoyBootstrapParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*pbdataplane.GetEnvoyBootstrapParamsRequest), variadicArgs...)
	})
	return _c
}

func (_c *DataplaneServiceClient_GetEnvoyBootstrapParams_Call) Return(_a0 *pbdataplane.GetEnvoyBootstrapParamsResponse, _a1 error) *DataplaneServiceClient_GetEnvoyBootstrapParams_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DataplaneServiceClient_GetEnvoyBootstrapParams_Call) RunAndReturn(run func(context.Context, *pbdataplane.GetEnvoyBootstrapParamsRequest, ...grpc.CallOption) (*pbdataplane.GetEnvoyBootstrapParamsResponse, error)) *DataplaneServiceClient_GetEnvoyBootstrapParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetSupportedDataplaneFeatures provides a mock function with given fields: ctx, in, opts
func (_m *DataplaneServiceClient) GetSupportedDataplaneFeatures(ctx context.Context, in *pbdataplane.GetSupportedDataplaneFeaturesRequest, opts ...grpc.CallOption) (*pbdataplane.GetSupportedDataplaneFeaturesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSupportedDataplaneFeatures")
	}

	var r0 *pbdataplane.GetSupportedDataplaneFeaturesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pbdataplane.GetSupportedDataplaneFeaturesRequest, ...grpc.CallOption) (*pbdataplane.GetSupportedDataplaneFeaturesResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pbdataplane.GetSupportedDataplaneFeaturesRequest, ...grpc.CallOption) *pbdataplane.GetSupportedDataplaneFeaturesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pbdataplane.GetSupportedDataplaneFeaturesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pbdataplane.GetSupportedDataplaneFeaturesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DataplaneServiceClient_GetSupportedDataplaneFeatures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSupportedDataplaneFeatures'
type DataplaneServiceClient_GetSupportedDataplaneFeatures_Call struct {
	*mock.Call
}

// GetSupportedDataplaneFeatures is a helper method to define mock.On call
//   - ctx context.Context
//   - in *pbdataplane.GetSupportedDataplaneFeaturesRequest
//   - opts ...grpc.CallOption
func (_e *DataplaneServiceClient_Expecter) GetSupportedDataplaneFeatures(ctx interface{}, in interface{}, opts ...interface{}) *DataplaneServiceClient_GetSupportedDataplaneFeatures_Call {
	return &DataplaneServiceClient_GetSupportedDataplaneFeatures_Call{Call: _e.mock.On("GetSupportedDataplaneFeatures",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *DataplaneServiceClient_GetSupportedDataplaneFeatures_Call) Run(run func(ctx context.Context, in *pbdataplane.GetSupportedDataplaneFeaturesRequest, opts ...grpc.CallOption)) *DataplaneServiceClient_GetSupportedDataplaneFeatures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*pbdataplane.GetSupportedDataplaneFeaturesRequest), variadicArgs...)
	})
	return _c
}

func (_c *DataplaneServiceClient_GetSupportedDataplaneFeatures_Call) Return(_a0 *pbdataplane.GetSupportedDataplaneFeaturesResponse, _a1 error) *DataplaneServiceClient_GetSupportedDataplaneFeatures_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DataplaneServiceClient_GetSupportedDataplaneFeatures_Call) RunAndReturn(run func(context.Context, *pbdataplane.GetSupportedDataplaneFeaturesRequest, ...grpc.CallOption) (*pbdataplane.GetSupportedDataplaneFeaturesResponse, error)) *DataplaneServiceClient_GetSupportedDataplaneFeatures_Call {
	_c.Call.Return(run)
	return _c
}

// NewDataplaneServiceClient creates a new instance of DataplaneServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDataplaneServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *DataplaneServiceClient {
	mock := &DataplaneServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
