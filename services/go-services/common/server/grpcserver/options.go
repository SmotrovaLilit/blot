package grpcserver

import "context"

type contextFactoryForRequestHandler func(ctx context.Context) context.Context

type contextFactoryForRequestHandlerOption struct {
	f contextFactoryForRequestHandler
}

func WithContextFactoryForRequestHandler(f contextFactoryForRequestHandler) Option {
	return contextFactoryForRequestHandlerOption{
		f: f,
	}
}

func (opt contextFactoryForRequestHandlerOption) apply(opts *options) {
	opts.contextFactory = opt.f
}

type options struct {
	contextFactory contextFactoryForRequestHandler
}

type Option interface {
	apply(*options)
}
