package internal

import (
	"github.com/quydmfl/payment-svc/internal/service"
	"github.com/quydmfl/payment-svc/proto"
	"github.com/ankorstore/yokai/fxgrpcserver"
	"go.uber.org/fx"
)

// Register is used to register the application dependencies.
func Register() fx.Option {
	return fx.Options(
		fxgrpcserver.AsGrpcServerService(service.NewExampleService, &proto.ExampleService_ServiceDesc),
	)
}
