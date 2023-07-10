package repository_factories

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	connectors_adapters "github.com/vm-championships-manager/backend-service/internal/connectors/adapters"
	repositories_adapters "github.com/vm-championships-manager/backend-service/internal/repositories/adapters"
	repositores_protocols "github.com/vm-championships-manager/backend-service/internal/repositories/protocols"
)

func NewUserRepository(ctx context.Context, awsCfg *aws.Config) repositores_protocols.UserRepository {
	dynamodb := connectors_adapters.NewDynamodbConnector(ctx, awsCfg)
	return repositories_adapters.NewUserRepository(*dynamodb.Dynamodb)
}
