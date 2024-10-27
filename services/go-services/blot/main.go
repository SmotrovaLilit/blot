package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"runtime/debug"

	"blot/internal/blot/ports"
	"blot/internal/blot/service"
	blotservicepb "blot/internal/common/gen-proto/blotservice/v1beta1"
	"blot/internal/common/logging"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	otelsdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"blot/internal/common/server"
)

func main() {
	ctx := context.Background()
	file, err := os.OpenFile("../../../deployment/logs/blot-api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	defer file.Close()
	logger := logging.NewLogger(file, false, slog.LevelDebug)
	buildInfo, _ := debug.ReadBuildInfo()
	pid := os.Getpid()
	ctx = logging.AppendCtx(
		ctx,
		slog.Int("pid", pid),
		slog.String("go_version", buildInfo.GoVersion),
	)
	slog.SetDefault(logger)
	app := service.NewApplication(ctx)
	tp, err := initOpenTelemetry(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			slog.ErrorContext(ctx, "Error tracer provider", "error", err)
		}
	}()
	server.RunGRPCServer(func(ctx context.Context) context.Context {
		ctx = logging.AppendCtx(ctx, slog.Group("build_info",
			slog.String("path", buildInfo.Main.Path),
			slog.String("go_version", buildInfo.GoVersion),
		))
		return ctx
	}, func(server *grpc.Server) {
		svc := ports.NewGrpcServer(app)
		reflection.Register(server)
		blotservicepb.RegisterBlotServiceServer(server, svc)
	})
}

func initOpenTelemetry(ctx context.Context) (*otelsdktrace.TracerProvider, error) {
	otlpEndpointHTTP := fmt.Sprintf("%s:%s", "127.0.0.1", "4318")
	insecureOpt := otlptracehttp.WithInsecure()
	endpointOptHTTP := otlptracehttp.WithEndpoint(otlpEndpointHTTP)
	exporterHTTP, err := otlptracehttp.New(ctx, insecureOpt, endpointOptHTTP)
	if err != nil {
		return nil, err
	}

	// Ensure default SDK resources and the required service name are set.
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("blot-api"),
			semconv.ServiceVersion("v1beta1"),
		),
	)
	if err != nil {
		log.Fatalf("failed to create otel resource: %v", err)
	}

	tp := otelsdktrace.NewTracerProvider(
		otelsdktrace.WithSampler(otelsdktrace.AlwaysSample()),
		otelsdktrace.WithBatcher(exporterHTTP),
		otelsdktrace.WithResource(r),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{}) // Required for propagating traceparent from UI
	slog.InfoContext(ctx, "OpenTelemetry initialized successfully with OTLP exporter over HTTP", "endpoint", otlpEndpointHTTP)
	return tp, nil
}
