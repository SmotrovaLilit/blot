package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"

	"blot/internal/common/logging"
	"blot/internal/common/safe"

	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

var tracer = otel.Tracer("grpc_handler")

type contextFactory func(ctx context.Context) context.Context

func RunGRPCServer(
	contextFactory contextFactory,
	registerServer func(server *grpc.Server),
) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	addr := fmt.Sprintf(":%s", port)
	RunGRPCServerOnAddr(addr, contextFactory, registerServer)
}

func newContextFactoryUnaryServerInterceptor(contextFactory contextFactory) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		ctx = contextFactory(ctx)
		return handler(ctx, req)
	}
}

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

	slog.InfoContext(ctx, "Executing unary request")

	resp, err = handler(ctx, req)

	if err != nil {
		slog.ErrorContext(ctx, "Failed to execute unary request", slog.Any("error", err))
	} else {
		slog.InfoContext(ctx, "Unary request executed successfully", slog.Any("grpc_response", resp))
	}
	return resp, err
}

func RunGRPCServerOnAddr(
	addr string,
	contextFactory contextFactory,
	registerServer func(server *grpc.Server),
) {
	grpcPanicRecoveryHandler := func(ctx context.Context, p interface{}) (err error) {
		safe.DefaultRecover(ctx, p)
		return status.Errorf(codes.Internal, "internal server error1")
	}
	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandlerContext(grpcPanicRecoveryHandler)),
			newContextFactoryUnaryServerInterceptor(contextFactory),
			loggingUnaryServerInterceptor,
		),
	)

	registerServer(grpcServer)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Starting gRPC server on %s\n", addr)
	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
