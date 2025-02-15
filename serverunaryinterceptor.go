// Copyright (c) 2022 Dmitry Tkachenko (tkachenkodmitryv@gmail.com)
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package care

import (
	"context"

	"google.golang.org/grpc"
)

type unaryServerInterceptor struct {
	interceptor *interceptor
}

func (s *unaryServerInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (
		interface{},
		error) {

		return s.interceptor.execute(
			ctx,
			info.FullMethod,
			req,
			func(c context.Context, r interface{}) (interface{}, error) {
				return handler(c, r)
			},
		)
	}
}

// NewServerUnaryInterceptor - makes a new unary server interceptor.
// There will be panic if options is an empty pointer.
func NewServerUnaryInterceptor(opts *Options) grpc.UnaryServerInterceptor {
	if opts == nil {
		panic("The options must not be provided as a nil-pointer.")
	}

	interceptor := unaryServerInterceptor{
		interceptor: newInterceptor(opts),
	}

	return interceptor.Unary()
}

func NewServerOption(opts *Options) grpc.ServerOption {
	if opts == nil {
		panic("The options must not be provided as a nil-pointer.")
	}

	interceptor := unaryServerInterceptor{
		interceptor: newInterceptor(opts),
	}

	return grpc.UnaryInterceptor(interceptor.Unary())
}
