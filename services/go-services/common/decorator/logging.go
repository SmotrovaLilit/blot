package decorator

import (
	"context"
	"log/slog"
)

type commandLoggingDecorator[C any] struct {
	base CommandHandler[C]
}

func (d commandLoggingDecorator[C]) Handle(ctx context.Context, cmd C) (err error) {
	handlerType := generateActionName(cmd)

	logger := slog.Default().With(
		slog.String("command", handlerType),
		slog.Any("command_body", cmd),
	)

	logger.Debug("Executing command")
	defer func() {
		if err == nil {
			logger.Info("Command executed successfully") // TODO it happens even when panic inside handler
		} else {
			logger.Error("Failed to execute command", slog.Any("error", err))
		}
	}()

	return d.base.Handle(ctx, cmd)
}

type queryLoggingDecorator[C any, R any] struct {
	base QueryHandler[C, R]
}

func (d queryLoggingDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err error) {
	logger := slog.Default().With(
		slog.String("query", generateActionName(cmd)),
		slog.Any("query_body", cmd),
	)

	logger.Debug("Executing query")
	defer func() {
		if err == nil {
			logger.Info("Query executed successfully")
			logger.Debug("Query executed successfully", "query_response", result)
		} else {
			logger.Error("Failed to execute query", slog.Any("error", err))
		}
	}()

	return d.base.Handle(ctx, cmd)
}
