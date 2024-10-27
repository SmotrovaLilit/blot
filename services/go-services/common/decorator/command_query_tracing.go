package decorator

import (
	"context"

	"go.opentelemetry.io/otel"
)

type commandTracingDecorator[C any] struct {
	base CommandHandler[C]
}

var tracer = otel.Tracer("application")

func (d commandTracingDecorator[C]) Handle(ctx context.Context, cmd C) (err error) {
	handlerType := generateActionName(cmd)
	ctx, span := tracer.Start(ctx, "application."+handlerType)
	defer span.End()
	return d.base.Handle(ctx, cmd)
}

type queryTracingDecorator[C any, R any] struct {
	base QueryHandler[C, R]
}

func (d queryTracingDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err error) {
	handlerType := generateActionName(cmd)
	ctx, span := tracer.Start(ctx, "application."+handlerType)
	defer span.End()

	return d.base.Handle(ctx, cmd)
}
