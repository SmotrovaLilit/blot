package grpcserver

import (
	"context"

	"google.golang.org/grpc"
)

func newContextFactoryUnaryServerInterceptor(contextFactory contextFactoryForRequestHandler) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if contextFactory != nil {
			ctx = contextFactory(ctx)
		}
		return handler(ctx, req)
	}
}
