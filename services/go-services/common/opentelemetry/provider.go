package opentelemetry

import (
	"context"
	"errors"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	otelsdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

type Config struct {
	Endpoint       string
	ServiceName    string
	ServiceVersion string
}

func Init(ctx context.Context, cfg Config) (*otelsdktrace.TracerProvider, error) {
	if cfg.Endpoint == "" {
		return nil, errors.New("OTLP endpoint is required")
	}
	if cfg.ServiceName == "" {
		return nil, errors.New("service name is required")
	}
	if cfg.ServiceVersion == "" {
		return nil, errors.New("service version is required")
	}
	otlpEndpointHTTP := cfg.Endpoint
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
			semconv.ServiceName(cfg.ServiceName),
			semconv.ServiceVersion(cfg.ServiceVersion),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create otel resource: %v", err)
	}

	tp := otelsdktrace.NewTracerProvider(
		otelsdktrace.WithSampler(otelsdktrace.AlwaysSample()),
		otelsdktrace.WithBatcher(exporterHTTP),
		otelsdktrace.WithResource(r),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{}) // Required for propagating trace parent from UI
	return tp, nil
}
