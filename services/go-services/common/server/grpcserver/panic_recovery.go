package grpcserver

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"blot/internal/common/safe"
)

func panicRecoveryHandler(ctx context.Context, p interface{}) (err error) {
	safe.DefaultRecover(ctx, p)
	return status.Errorf(codes.Internal, "internal server error")
}
