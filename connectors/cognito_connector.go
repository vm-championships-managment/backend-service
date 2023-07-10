package connectors

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/vm-championships-manager/backend-service/config"
	utils_protocols "github.com/vm-championships-manager/backend-service/internal/utils/protocols"
)

type CognitoConnector interface {
	Exec(utils_protocols.CommandGenericInput[struct{ email, password string }]) utils_protocols.CommandGenericOutput[*string]
}

type CognitoConnectorAdapter struct {
	client *cognitoidentityprovider.Client
	// credentials *config.Credentials
}

func NewCognitoAdapter(ctx context.Context, awsDefaultCfg *aws.Config) *CognitoConnectorAdapter {
	cOpts := cognitoidentityprovider.EndpointResolverFunc(func(region string, options cognitoidentityprovider.EndpointResolverOptions) (aws.Endpoint, error) {
		return getEndpointResolver(config.AwsEndpoint, awsDefaultCfg.Region)
	})
	client := cognitoidentityprovider.NewFromConfig(*awsDefaultCfg, cognitoidentityprovider.WithEndpointResolver(cOpts))

	return &CognitoConnectorAdapter{
		client: client,
	}
	// cca.credentials.LoadConfig()

	// done := make(chan bool)
	// ctxTimeout, ctxCancel := context.WithTimeout(context.Background(), time.Second)
	// defer ctxCancel()
	// go func() {
	// 	sess := pkg.MustDouble(session.NewSession(&aws.Config{
	// 		Region:      aws.String(cca.credentials.Aws.Region),
	// 		Credentials: credentials.NewStaticCredentials(cca.credentials.Aws.AccessKeyID, cca.credentials.Aws.SecretAccessKey, ""),
	// 		Endpoint:    aws.String(cca.credentials.Aws.Dynamodb.Endpoint),
	// 	}))
	// 	cca.client = cognitoidentityprovider.New(sess)
	// 	done <- true
	// }()

	// select {
	// case <-ctxTimeout.Done():
	// 	panic(fmt.Sprintf("COGNITO CONNECTOR: %s", ctxTimeout.Err()))
	// case <-done:
	// 	return cca
	// }
}

// func (cca CognitoConnectorAdapter) Exec(in utils.Input[struct{ email, password string }]) (out utils.Output[*string]) {
// 	switch in.Action {
// 	case "USER/CREATE":
// 		cca.createUser(in.Payload)
// 	default:
// 		out.Error = internal_errors.NewInvalidCommandError("COGNITO_CONNECTOR")
// 	}

// 	return
// }

// func (cca CognitoConnectorAdapter) createUser(args struct{ email, password string }) {
// 	var poolId string

// 	for _, role := range cca.credentials.Aws.Cognito.Roles {
// 		if role.Name == "users" {
// 			poolId = role.PoolId
// 		}
// 	}

// 	input := &cognitoidentityprovider.AdminCreateUserInput{
// 		UserPoolId:        aws.String(poolId),
// 		Username:          aws.String(args.email),
// 		TemporaryPassword: aws.String(args.password),
// 		UserAttributes: []*cognitoidentityprovider.AttributeType{
// 			{
// 				Name:  aws.String("email"),
// 				Value: aws.String(args.email),
// 			},
// 		},
// 	}

// 	pkg.MustDouble(cca.client.AdminCreateUser(input))
// }
