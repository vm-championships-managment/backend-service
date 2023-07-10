package controllers_factories

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	controllers_adapters "github.com/vm-championships-manager/backend-service/internal/controllers/adapters"
	controllers_protocols "github.com/vm-championships-manager/backend-service/internal/controllers/protocols"
	services_factories "github.com/vm-championships-manager/backend-service/internal/services/factories"
	utils_adapters "github.com/vm-championships-manager/backend-service/internal/utils/adapters"
)

func NewUserController(ctx context.Context, defaultCfg *aws.Config) controllers_protocols.Controller {
	usrSvc := services_factories.NewUserService(ctx, defaultCfg)
	return controllers_adapters.NewUserController(usrSvc, utils_adapters.NewLoggerAdapter())
}
