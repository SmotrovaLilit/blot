package logging

import (
	"io"
	"log/slog"
)

func NewLogger(w io.Writer, humanReadable bool, level slog.Level) *slog.Logger {
	var base slog.Handler
	if humanReadable {
		base = slog.NewTextHandler(w, &slog.HandlerOptions{
			Level: level,
		})
	} else {
		base = slog.NewJSONHandler(w, &slog.HandlerOptions{
			Level: level,
		})
	}
	h := &ContextHandler{Handler: base}
	return slog.New(&OpenTelemetryHandler{Handler: h})
}
