package logging

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"
)

type OpenTelemetryHandler struct {
	slog.Handler
}

// Handle adds trace fields to every log record.
func (h OpenTelemetryHandler) Handle(ctx context.Context, r slog.Record) error {
	span := trace.SpanFromContext(ctx)
	r.Add("trace_id", slog.StringValue(span.SpanContext().TraceID().String()))
	r.Add("span_id", slog.StringValue(span.SpanContext().SpanID().String()))
	return h.Handler.Handle(ctx, r)
}
