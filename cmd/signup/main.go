package main

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vm-championships-manager/backend-service/connectors"
	controllers_factories "github.com/vm-championships-manager/backend-service/internal/controllers/factories"
	controllers_protocols "github.com/vm-championships-manager/backend-service/internal/controllers/protocols"
	errors_adapters "github.com/vm-championships-manager/backend-service/internal/errors/adapters"
	utils_adapters "github.com/vm-championships-manager/backend-service/internal/utils/adapters"
)

func main() {
	lambda.Start(router)
}

func router(req events.APIGatewayProxyRequest) (cOut controllers_protocols.ControllerOutput, err error) {
	defer func() {
		if err := recover(); err != nil {
			log := utils_adapters.NewLoggerAdapter()
			log.Metadata(map[string]interface{}{
				"stacktrace": string(debug.Stack()),
			}).Error("[Signup.main] panic occured")
			intServErr := errors_adapters.NewInternalServerError()

			cOut.StatusCode = intServErr.GetErrorInfos().Code
			cOut.Body = string(
				fmt.Sprintf(
					`{"error": {"name": "%s", "message": "%s"}}`,
					intServErr.GetErrorInfos().Name,
					intServErr.GetErrorInfos().Message,
				))
		}
	}()

	ctx, ctxCancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer ctxCancel()

	awsCfg := connectors.LoadAwsConfig(ctx)
	uc := controllers_factories.NewUserController(ctx, awsCfg)

	payload := make(map[string]interface{})
	if e := json.Unmarshal([]byte(req.Body), &payload); e != nil {
		panic(e)
	}

	cOut = uc.Handle(ctx, controllers_protocols.ControllerInput{
		Action: controllers_protocols.UserControllerActionCreate,
		Payload: controllers_protocols.ControllerPayload{
			Body: payload,
		},
	})

	return
}
