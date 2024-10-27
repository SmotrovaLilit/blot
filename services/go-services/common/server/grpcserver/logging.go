package grpcserver

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"blot/internal/common/logging"
)

func loggingUnaryServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	ctx, span := tracer.Start(ctx, info.FullMethod)
	defer span.End()
	md, _ := metadata.FromIncomingContext(ctx)

	ctx = logging.AppendCtx(
		ctx,
		slog.String("grpc_method", info.FullMethod),
		slog.Any("grpc_metadata", md),
		slog.Any("grpc_request", req),
	)

	slog.DebugContext(ctx, "Executing unary request")

	resp, err = handler(ctx, req)

	if err != nil {
		slog.ErrorContext(ctx, "Failed to execute unary request", slog.Any("error", err))
	} else {
		slog.InfoContext(ctx, "Unary request executed successfully", slog.Any("grpc_response", resp))
	}
	return resp, err
}
