package grpcserver

import (
	"fmt"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func RunServerOnAddr(
	addr string,
	registerServer func(server *grpc.Server),
	opts ...Option,
) error {
	resultOptions := options{}
	for _, opt := range opts {
		opt.apply(&resultOptions)
	}
	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandlerContext(panicRecoveryHandler)),
			newContextFactoryUnaryServerInterceptor(resultOptions.contextFactory),
			tracingUnaryServerInterceptor,
			loggingUnaryServerInterceptor,
		),
	)

	registerServer(grpcServer)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen %s: %w", addr, err)
	}

	err = grpcServer.Serve(listen)
	if err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}
	return nil
}
