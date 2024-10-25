package safe

import (
	"context"
	"fmt"
	"log/slog"
	"runtime/debug"
)

// GoContextWithRecover starts a goroutine with a recover function.
// The recover function is called when the goroutine panics.
// The recover function is called with the context and the panic value.
func GoContextWithRecover(ctx context.Context, goroutine func(), customRecover func(context.Context, interface{})) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				customRecover(ctx, r)
			}
		}()
		goroutine()
	}()
}

// DefaultRecover is the default recover function.
// It logs the panic with the stack trace.
// Stack trace will be in the field "stacktrace" base64 encoded.
// To discover errors you can use the following command:
// echo "base64 encoded stacktrace" | base64 -d
func DefaultRecover(ctx context.Context, p interface{}) {
	slog.ErrorContext(ctx, fmt.Sprintf("panic recovered: %v", p), "stacktrace", debug.Stack(), "error", p)
}

// GoContext starts a goroutine with the default recover function.
// The default recover function logs the panic with the stack trace.
// The recover function is called with the context and the panic value.
func GoContext(ctx context.Context, goroutine func()) {
	GoContextWithRecover(ctx, goroutine, DefaultRecover)
}
