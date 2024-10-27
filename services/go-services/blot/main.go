package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"blot/internal/common/server/grpcserver"

	"blot/internal/blot/ports"
	"blot/internal/blot/service"
	blotservicepb "blot/internal/common/gen-proto/blotservice/v1beta1"
	"blot/internal/common/logging"
	"blot/internal/common/opentelemetry"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()
	file, err := os.OpenFile("../../../deployment/logs/blot-api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	defer file.Close()
	logger := logging.NewLogger(file, false, slog.LevelDebug)
	slog.SetDefault(logger)
	application := service.NewApplication(ctx)
	ctx = application.AppendCtxWithApplicationLoggingFields(ctx)
	cfg := opentelemetry.Config{
		Endpoint:       fmt.Sprintf("%s:%s", "127.0.0.1", "4318"),
		ServiceName:    application.Info.Name,
		ServiceVersion: application.Info.Version,
	}
	tp, err := opentelemetry.Init(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to initialize openTelemetry: %v", err)
	} else {
		slog.InfoContext(ctx, "openTelemetry initialized successfully with OLTP exporter over HTTP", "cfg", cfg)
	}
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			slog.ErrorContext(ctx, "Error tracer provider", "error", err)
		}
	}()

	addr := ":8081"
	err = grpcserver.RunServerOnAddr(
		addr,
		func(server *grpc.Server) {
			svc := ports.NewGrpcServer(application)
			reflection.Register(server)
			blotservicepb.RegisterBlotServiceServer(server, svc)
		},
		grpcserver.WithContextFactoryForRequestHandler(application.AppendCtxWithApplicationLoggingFields),
	)
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
	fmt.Printf("Starting gRPC server on %s\n", addr)
}
