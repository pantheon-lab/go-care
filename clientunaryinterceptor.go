// Copyright (c) 2022 Dmitry Tkachenko (tkachenkodmitryv@gmail.com)
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package care

import (
	"context"
	"reflect"

	"google.golang.org/grpc"
)

type unaryClientInterceptor struct {
	interceptor *interceptor
}

func (s *unaryClientInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption) error {

		realComputation := false

		resp, err := s.interceptor.execute(
			ctx,
			method,
			req,
			func(c context.Context, r interface{}) (interface{}, error) {
				e := invoker(c, method, r, reply, cc, opts...)
				realComputation = true
				return reply, e
			},
		)

		if !realComputation && resp != nil {
			replyVal := reflect.ValueOf(reply).Elem()
			respVal := reflect.ValueOf(resp).Elem()
			tmp := replyVal.Interface()

			replyVal.Set(respVal)
			respVal.Set(reflect.ValueOf(tmp))
		}

		return err
	}
}

// NewClientUnaryInterceptor makes a new unary client interceptor.
// There will be panic if options is an empty pointer.
func NewClientUnaryInterceptor(opts *Options) grpc.UnaryClientInterceptor {
	if opts == nil {
		panic("The options must not be provided as a nil-pointer.")
	}

	interceptor := unaryClientInterceptor{
		interceptor: newInterceptor(opts),
	}

	return interceptor.Unary()
}

func NewClientDialOption(opts *Options) grpc.DialOption {
	if opts == nil {
		panic("The options must not be provided as a nil-pointer.")
	}

	interceptor := unaryClientInterceptor{
		interceptor: newInterceptor(opts),
	}

	return grpc.WithUnaryInterceptor(interceptor.Unary())
}
