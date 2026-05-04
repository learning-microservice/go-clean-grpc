package log

import (
	"fmt"
	"log/slog"
	"strings"
)

// Option -.
type Option func(*config)

type config struct {
	option slog.HandlerOptions
	attrs  []any
}

// Level -.
func Level(level string) Option {
	return func(c *config) {
		switch strings.ToLower(level) {
		case "error":
			c.option.Level = slog.LevelError
		case "warn":
			c.option.Level = slog.LevelWarn
		case "info":
			c.option.Level = slog.LevelInfo
		case "debug":
			c.option.Level = slog.LevelDebug
		default:
			panic(fmt.Sprintf("invalid log level: %s", level))
		}
	}
}

// GlobalAttrs -.
func GlobalAttrs(attrs ...any) Option {
	return func(c *config) {
		c.attrs = attrs
	}
}
