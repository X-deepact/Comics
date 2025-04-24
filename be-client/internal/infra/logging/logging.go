package logging

import (
	"be-client/config"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"
)

var counter uint64

type LogEntry struct {
	Level      string    `json:"level"`
	Timestamp  time.Time `json:"timestamp"`
	Message    string    `json:"message"`
	Error      string    `json:"error,omitempty"`
	Caller     string    `json:"caller,omitempty"`
	Additional any       `json:"additional,omitempty"`
}

type logHandler struct {
	slog.Handler
	additionalAttrs []slog.Attr
}

func InitLogger(cf config.Config) *slog.Logger {
	cfg := cf.Log
	programLevel := new(slog.LevelVar)
	programLevel.Set(cfg.LogLevel)

	handlerOptions := &slog.HandlerOptions{
		AddSource: cfg.AddSource,
		Level:     programLevel,
	}

	var baseHandler slog.Handler
	if cfg.JSONOutput {
		baseHandler = slog.NewJSONHandler(os.Stdout, handlerOptions)
	} else {
		baseHandler = slog.NewTextHandler(os.Stdout, handlerOptions)
	}

	handler := &logHandler{
		Handler: baseHandler,
		additionalAttrs: []slog.Attr{
			slog.String("environment", cf.Profile.Env),
			slog.Time("boot_time", time.Now()),
		},
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)

	return logger
}

func (h *logHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, attr := range h.additionalAttrs {
		r.AddAttrs(attr)
	}

	return h.Handler.Handle(ctx, r)
}

func (h *logHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &logHandler{
		Handler:         h.Handler.WithAttrs(attrs),
		additionalAttrs: h.additionalAttrs,
	}
}

func ErrAttr(err error) slog.Attr {
	return slog.Any("error", err)
}

func LogWithContext(ctx context.Context, level slog.Level, msg string, args ...any) {
	if traceID := ctx.Value("trace_id"); traceID != nil {
		args = append(args, slog.String("trace_id", fmt.Sprint(traceID)))
	}

	slog.Log(ctx, level, msg, args...)
}

func LogError(ctx context.Context, err error, msg string, args ...any) {
	if err != nil {
		args = append(args, ErrAttr(err))
	}
	LogWithContext(ctx, slog.LevelError, msg, args...)
}

func ToJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("error marshaling to JSON: %v", err)
	}
	return string(b)
}
