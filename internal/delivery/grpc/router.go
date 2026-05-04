package grpc

import (
	"go-clean-grpc/api/v1/v1connect"
	"go-clean-grpc/internal/delivery/grpc/v1/auth"
	"go-clean-grpc/internal/registry"
	"go-clean-grpc/pkg/grpcserver"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
	"connectrpc.com/validate"
)

// SetupRouter -.
func SetupRouter(app *grpcserver.Server, reg *registry.Registry) {
	// setup interceptors
	interceptors := []connect.HandlerOption{
		// Validation via Protovalidate is almost always recommended
		connect.WithInterceptors(validate.NewInterceptor()),
	}

	// grpc.health.v1.Health/Check
	app.Handle(grpchealth.NewHandler(grpchealth.NewStaticChecker(
		v1connect.AuthServiceName,
	)))

	// setup auth service
	// grpc.v1.AuthService/Login
	app.Handle(v1connect.NewAuthServiceHandler(
		auth.NewLoginHandler(reg.UsecaseSet.UserLogin),
		interceptors...,
	))
}
