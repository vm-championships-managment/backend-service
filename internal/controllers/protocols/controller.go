package controllers_protocols

import (
	"context"

	utils_protocols "github.com/vm-championships-manager/backend-service/internal/utils/protocols"
)

const (
	UserControllerActionCreate = "USER/CREATE"
)

type ControllerPayload struct {
	Body    map[string]interface{}
	Headers map[string]interface{}
	Params  map[string]interface{}
}

type ControllerInput = utils_protocols.CommandGenericInput[ControllerPayload]
type ControllerOutput struct {
	StatusCode uint16 `json:"statusCode"`
	Body       string `json:"body"`
}

type Controller interface {
	Handle(ctx context.Context, in ControllerInput) ControllerOutput
}
