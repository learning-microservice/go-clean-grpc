package log

import (
	"io"
	"log/slog"
)

// New -.
func New(writer io.Writer, opts ...Option) *slog.Logger {
	cfg := config{
		option: slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: false,
		},
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	writer = &prettyJSONWriter{out: writer, indent: "  "}

	logger := slog.New(slog.NewJSONHandler(writer, &cfg.option))
	if len(cfg.attrs) > 0 {
		logger = logger.With(cfg.attrs...)
	}

	return logger
}

/*

// Debug -.
func (l *logger) Debug(ctx context.Context, message any, args ...any) {
	l.log(ctx, slog.LevelDebug, message, args...)
}

// Info -.
func (l *logger) Info(ctx context.Context, message string, args ...any) {
	l.log(ctx, slog.LevelInfo, message, args...)
}

// Warn -.
func (l *logger) Warn(ctx context.Context, message string, args ...any) {
	l.log(ctx, slog.LevelWarn, message, args...)
}

// Error -.
func (l *logger) Error(ctx context.Context, message any, args ...any) {
	l.log(ctx, slog.LevelError, message, args...)
}

func (l *logger) log(ctx context.Context, level slog.Level, message any, args ...any) {
	switch msg := message.(type) {
	case error:
		l.internal.Log(ctx, level, msg.Error(), args...)
	case string:
		l.internal.Log(ctx, level, msg, args...)
	default:
		l.internal.Log(ctx, level, fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}
*/
