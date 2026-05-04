package grpcserver

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	_defaultAddr              = ":8080"
	_defaultReadHeaderTimeout = 5 * time.Second
	_defaultShutdownTimeout   = 2 * time.Second
)

// Server -.
type Server struct {
	engine            *http.ServeMux
	address           string
	readHeaderTimeout time.Duration
	shutdownTimeout   time.Duration
	logger            *slog.Logger
}

// New -.
func New(opts ...Option) *Server {
	// setup server
	s := Server{
		engine:            http.NewServeMux(),
		address:           _defaultAddr,
		readHeaderTimeout: _defaultReadHeaderTimeout,
		shutdownTimeout:   _defaultShutdownTimeout,
		logger:            slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	// bind options
	for _, opt := range opts {
		opt(&s)
	}

	return &s
}

func (s *Server) Handle(path string, handler http.Handler) *Server {
	s.engine.Handle(path, handler)
	return s
}

// ListenAndServe -.
func (s *Server) ListenAndServe() error {
	// setup protocols
	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetUnencryptedHTTP2(true) // Use h2c so we can serve HTTP/2 without TLS.

	// setup server
	server := http.Server{
		Addr:              _defaultAddr,
		Handler:           s.engine,
		Protocols:         protocols,
		ReadHeaderTimeout: s.readHeaderTimeout,
		ErrorLog:          slog.NewLogLogger(s.logger.Handler(), slog.LevelError),
	}

	sigCtx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,    // ctrl+c
		syscall.SIGTERM, // kill command
		syscall.SIGHUP,  // hung up
	)
	defer stop()

	g, ctx := errgroup.WithContext(sigCtx)
	g.SetLimit(2)

	g.Go(func() error {
		s.logger.Info("starting grpc server...", "address", s.address)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	g.Go(func() error {
		<-ctx.Done()

		s.logger.Info("starting grpc server shutting down...", "address", s.address)
		ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		s.logger.Error("grpc server error", "error", err)
		return err
	}

	s.logger.Info("shutdown grpc server completed")
	return nil
}
