// Code generated by thriftrw-plugin-yarpc
// @generated

package readonlystoreclient

import (
	context "context"
	wire "go.uber.org/thriftrw/wire"
	yarpc "go.uber.org/yarpc"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	atomic "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/atomic"
	baseserviceclient "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/common/baseserviceclient"
	reflect "reflect"
)

// Interface is a client for the ReadOnlyStore service.
type Interface interface {
	baseserviceclient.Interface

	Integer(
		ctx context.Context,
		Key *string,
		opts ...yarpc.CallOption,
	) (int64, error)
}

// New builds a new client for the ReadOnlyStore service.
//
// 	client := readonlystoreclient.New(dispatcher.ClientConfig("readonlystore"))
func New(c transport.ClientConfig, opts ...thrift.ClientOption) Interface {
	return client{
		c: thrift.New(thrift.Config{
			Service:      "ReadOnlyStore",
			ClientConfig: c,
		}, opts...),
		nwc: thrift.NewNoWire(thrift.Config{
			Service:      "ReadOnlyStore",
			ClientConfig: c,
		}, opts...),

		Interface: baseserviceclient.New(
			c,
			append(
				opts,
				thrift.Named("ReadOnlyStore"),
			)...,
		),
	}
}

func init() {
	yarpc.RegisterClientBuilder(
		func(c transport.ClientConfig, f reflect.StructField) Interface {
			return New(c, thrift.ClientBuilderOptions(c, f)...)
		},
	)
}

type client struct {
	baseserviceclient.Interface

	c   thrift.Client
	nwc thrift.NoWireClient
}

func (c client) Integer(
	ctx context.Context,
	_Key *string,
	opts ...yarpc.CallOption,
) (success int64, err error) {

	var result atomic.ReadOnlyStore_Integer_Result
	args := atomic.ReadOnlyStore_Integer_Helper.Args(_Key)

	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	success, err = atomic.ReadOnlyStore_Integer_Helper.UnwrapResponse(&result)
	return
}
