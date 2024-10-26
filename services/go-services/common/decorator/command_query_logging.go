package decorator

import (
	"context"
	"log/slog"

	"blot/internal/common/logging"

	"go.opentelemetry.io/otel"
)

type commandLoggingDecorator[C any] struct {
	base CommandHandler[C]
}

var tracer = otel.Tracer("application")

func (d commandLoggingDecorator[C]) Handle(ctx context.Context, cmd C) (err error) {
	handlerType := generateActionName(cmd)
	ctx, span := tracer.Start(ctx, "application."+handlerType)
	defer span.End()
	ctx = logging.AppendCtx(ctx, slog.String("command", handlerType), slog.Any("command_body", cmd))

	slog.DebugContext(ctx, "Executing command")
	defer func() {
		if err == nil {
			slog.InfoContext(ctx, "Command executed successfully") // TODO it happens even when panic inside handler
		} else {
			slog.ErrorContext(ctx, "Failed to execute command", slog.Any("error", err))
		}
	}()

	return d.base.Handle(ctx, cmd)
}

type queryLoggingDecorator[C any, R any] struct {
	base QueryHandler[C, R]
}

func (d queryLoggingDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err error) {
	handlerType := generateActionName(cmd)
	ctx, span := tracer.Start(ctx, "application."+handlerType)
	defer span.End()
	ctx = logging.AppendCtx(ctx, slog.String("query", handlerType))
	ctx = logging.AppendCtx(ctx, slog.Any("query_body", cmd))

	slog.DebugContext(ctx, "Executing query")
	defer func() {
		if err == nil {
			slog.InfoContext(ctx, "Query executed successfully")
			slog.DebugContext(ctx, "Query executed successfully", "query_response", result)
		} else {
			slog.ErrorContext(ctx, "Failed to execute query", slog.Any("error", err))
		}
	}()

	return d.base.Handle(ctx, cmd)
}
