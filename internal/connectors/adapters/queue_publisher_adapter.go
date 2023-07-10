package connectors_adapters

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/vm-championships-manager/backend-service/config"
	connectors_protocols "github.com/vm-championships-manager/backend-service/internal/connectors/protocols"
)

type QueuePublisherAdapter struct {
	queueName string
	client    *sqs.Client
}

func NewQueuePublisher(parentCtx context.Context, queueName string, awsDefaultCfg *aws.Config) connectors_protocols.QueuePublisher {
	ctx, cancelCtx := context.WithTimeout(parentCtx, time.Second)
	defer cancelCtx()

	done := make(chan QueuePublisherAdapter, 1)
	go func() {
		cOpts := sqs.EndpointResolverFunc(func(region string, options sqs.EndpointResolverOptions) (aws.Endpoint, error) {
			return getEndpointResolver(config.AwsEndpoint, awsDefaultCfg.Region)
		})
		client := sqs.NewFromConfig(*awsDefaultCfg, sqs.WithEndpointResolver(cOpts))

		done <- QueuePublisherAdapter{
			client:    client,
			queueName: queueName,
		}
	}()

	select {
	case <-ctx.Done():
		panic(ctx.Err())
	case qpa := <-done:
		return &qpa
	}
}

func getEndpointResolver(endpoint string, region string) (aws.Endpoint, error) {
	if len(endpoint) > 0 {
		return aws.Endpoint{
			PartitionID:   "aws",
			URL:           endpoint,
			SigningRegion: region,
		}, nil
	}
	return aws.Endpoint{}, &aws.EndpointNotFoundError{}
}

func (qpa *QueuePublisherAdapter) Send(message string) {}
