package connectors_adapters

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/vm-championships-manager/backend-service/config"
)

type DynamodbConnectorAdapter struct {
	Dynamodb *dynamodb.Client
}

func NewDynamodbConnector(parentCtx context.Context, awsDefaultCfg *aws.Config) *DynamodbConnectorAdapter {
	ctx, cancel := context.WithTimeout(parentCtx, time.Second)
	defer cancel()

	done := make(chan bool, 1)
	var dca *DynamodbConnectorAdapter = &DynamodbConnectorAdapter{}
	go func(dca *DynamodbConnectorAdapter) {
		dOpts := dynamodb.EndpointResolverFunc(func(region string, options dynamodb.EndpointResolverOptions) (aws.Endpoint, error) {
			return getEndpointResolver(config.AwsEndpoint, awsDefaultCfg.Region)
		})

		dca.Dynamodb = dynamodb.NewFromConfig(
			*awsDefaultCfg,
			dynamodb.WithEndpointResolver(dOpts),
		)

		done <- true
	}(dca)

	select {
	case <-ctx.Done():
		panic(ctx.Err())
	case <-done:
		return dca
	}
}
