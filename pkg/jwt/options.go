package jwt

import (
	"time"
)

// Option -.
type Option func(*Manager)

// NowFunc -.
func NowFunc(nowFunc func() time.Time) Option {
	return func(m *Manager) {
		m.now = nowFunc
	}
}

// Duration -.
func Duration(duration time.Duration) Option {
	return func(m *Manager) {
		m.duration = duration
	}
}
