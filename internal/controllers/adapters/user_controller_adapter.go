package controllers_adapters

import (
	"context"
	"encoding/json"
	"fmt"

	controllers_protocols "github.com/vm-championships-manager/backend-service/internal/controllers/protocols"
	"github.com/vm-championships-manager/backend-service/internal/entities"
	errors_adapters "github.com/vm-championships-manager/backend-service/internal/errors/adapters"
	errors_protocols "github.com/vm-championships-manager/backend-service/internal/errors/protocols"
	pkg_adapter "github.com/vm-championships-manager/backend-service/internal/pkg/adapters"
	pkg_protocols "github.com/vm-championships-manager/backend-service/internal/pkg/protocols"
	services_protocols "github.com/vm-championships-manager/backend-service/internal/services/protocols"
	utils_protocols "github.com/vm-championships-manager/backend-service/internal/utils/protocols"
)

type input = controllers_protocols.ControllerInput
type output = controllers_protocols.ControllerOutput

type UserControllerAdapter struct {
	svc          services_protocols.UserService
	log          utils_protocols.Log
	errorHandler pkg_protocols.ErrorHandler
}

func NewUserController(svc services_protocols.UserService, log utils_protocols.Log) controllers_protocols.Controller {
	return &UserControllerAdapter{
		svc:          svc,
		log:          log,
		errorHandler: pkg_adapter.NewPanicIfErrorOccur(log),
	}
}

func (c *UserControllerAdapter) Handle(parentCtx context.Context, in input) output {
	ctx, ctxCancel := context.WithCancel(parentCtx)
	defer ctxCancel()

	done := make(chan output, 1)

	go func() {
		var out output
		switch in.Action {
		case controllers_protocols.UserControllerActionCreate:
			out = c.create(ctx, in.Payload.Body)
		default:
			c.log.Metadata(map[string]interface{}{"input": in}).Error(
				"[UserController.Exec] unrecognized command",
			)
			out = c.fmtOutputErrors(errors_adapters.NewInternalServerError())
		}

		done <- out
	}()

	select {
	case <-ctx.Done():
		c.log.Metadata(map[string]interface{}{"input": in}).Error("[UserController.Exec] context done")
		panic("[UserController.Exec] context done")
	case out := <-done:
		return out
	}
}

func (c *UserControllerAdapter) fmtOutputErrors(err errors_protocols.CustomError) controllers_protocols.ControllerOutput {
	return controllers_protocols.ControllerOutput{
		StatusCode: err.GetErrorInfos().Code,
		Body: fmt.Sprintf(
			"{\"error\": {\"name\": \"%s\", \"message\": \"%s\"}}",
			err.GetErrorInfos().Name,
			err.GetErrorInfos().Message,
		)}
}

func (c *UserControllerAdapter) create(ctx context.Context, payload map[string]interface{}) controllers_protocols.ControllerOutput {
	logErrorParams := func(msg string) (string, map[string]interface{}) {
		return fmt.Sprintf("[UserController] %s", msg), map[string]interface{}{"payload": payload}
	}

	data := c.errorHandler.Double(
		json.Marshal(payload),
	)(logErrorParams("marshal json got an error")).([]byte)

	u := entities.User{}
	c.errorHandler.Single(json.Unmarshal(data, &u))(logErrorParams("unmarshal json got an error"))

	if entityErr := u.Validate(); entityErr != nil {
		c.log.Metadata(
			map[string]interface{}{
				"error": entityErr,
				"metadata": map[string]interface{}{
					"user": u,
				},
			},
		).Error("[UserController] validate entity error")
		return c.fmtOutputErrors(entityErr)
	}

	out := c.svc.Exec(ctx, services_protocols.UserServiceInput{
		Action:  services_protocols.ActionUserCreate,
		Payload: struct{ entities.User }{u},
	})

	if out.Error != nil {
		return c.fmtOutputErrors(out.Error)
	}

	data = c.errorHandler.Double(json.Marshal(out.Result.UserDto))(
		logErrorParams("marshal json got an error"),
	).([]byte)

	return controllers_protocols.ControllerOutput{
		StatusCode: 201,
		Body:       string(data),
	}
}
