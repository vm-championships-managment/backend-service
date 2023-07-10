package services_factories

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	repositories_factories "github.com/vm-championships-manager/backend-service/internal/repositories/factories"
	services_adapters "github.com/vm-championships-manager/backend-service/internal/services/adapters"
	services_protocols "github.com/vm-championships-manager/backend-service/internal/services/protocols"
	utils_adapters "github.com/vm-championships-manager/backend-service/internal/utils/adapters"
)

func NewUserService(ctx context.Context, defaultCfg *aws.Config) services_protocols.UserService {
	return services_adapters.NewUserService(
		repositories_factories.NewUserRepository(ctx, defaultCfg),
		utils_adapters.NewLoggerAdapter(),
	)
}
