// Code generated by Kitex v1.4.2. DO NOT EDIT.

package echo

import (
	"github.com/cloudwego/examples/kitex_gen/pbapi"
	"github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler pbapi.Echo, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
