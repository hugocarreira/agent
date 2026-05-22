package log

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

type plainHandler struct{}

func (h *plainHandler) Enabled(ctx context.Context, l slog.Level) bool {
	return l >= slog.LevelInfo
}

func (h *plainHandler) Handle(ctx context.Context, r slog.Record) error {
	fmt.Fprintln(os.Stdout, r.Message)
	return nil
}

func (h *plainHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *plainHandler) WithGroup(name string) slog.Handler {
	return h
}

var L = slog.New(&plainHandler{})
