package connectors

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func LoadAwsConfig(ctx context.Context) *aws.Config {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Panicln("Failed to load AWS configuration:", err)
	}

	return &cfg
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
