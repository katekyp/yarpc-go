// Code generated by thriftrw-plugin-yarpc
// @generated

package emptyservicefx

import (
	fx "go.uber.org/fx"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	emptyserviceserver "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/common/emptyserviceserver"
)

// ServerParams defines the dependencies for the EmptyService server.
type ServerParams struct {
	fx.In

	Handler emptyserviceserver.Interface
}

// ServerResult defines the output of EmptyService server module. It provides the
// procedures of a EmptyService handler to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group. Dig 1.2 or newer
// must be used for this feature to work.
type ServerResult struct {
	fx.Out

	Procedures []transport.Procedure `group:"yarpcfx"`
}

// Server provides procedures for EmptyService to an Fx application. It expects a
// emptyservicefx.Interface to be present in the container.
//
// 	fx.Provide(
// 		func(h *MyEmptyServiceHandler) emptyserviceserver.Interface {
// 			return h
// 		},
// 		emptyservicefx.Server(),
// 	)
func Server(opts ...thrift.RegisterOption) interface{} {
	return func(p ServerParams) ServerResult {
		procedures := emptyserviceserver.New(p.Handler, opts...)
		return ServerResult{Procedures: procedures}
	}
}
