package grpcserver

import (
	"context"

	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
)

var tracer = otel.Tracer("grpc_handler")

func tracingUnaryServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	ctx, span := tracer.Start(ctx, info.FullMethod)
	defer span.End()

	return handler(ctx, req)
}
