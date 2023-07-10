package services_adapters

import (
	"context"
	"fmt"

	"github.com/vm-championships-manager/backend-service/internal/dtos"
	"github.com/vm-championships-manager/backend-service/internal/entities"
	errors_adapters "github.com/vm-championships-manager/backend-service/internal/errors/adapters"
	repositories_protocols "github.com/vm-championships-manager/backend-service/internal/repositories/protocols"
	services_protocols "github.com/vm-championships-manager/backend-service/internal/services/protocols"
	utils_protocols "github.com/vm-championships-manager/backend-service/internal/utils/protocols"
)

type input = services_protocols.UserServiceInput
type output = services_protocols.UserServiceOutput

type UsrSvcInput struct {
	entities.User
}

type UsrSvcOutput struct {
	*dtos.UserDto
}

type userRepository = repositories_protocols.UserRepository

type UserSvc struct {
	usrRepo userRepository
	log     utils_protocols.Log
}

func NewUserService(usrRepo repositories_protocols.UserRepository, log utils_protocols.Log) *UserSvc {
	return &UserSvc{
		usrRepo: usrRepo,
		log:     log,
	}
}

func (svc *UserSvc) Exec(parentCtx context.Context, in input) output {
	ctx, cancelCtx := context.WithCancel(parentCtx)
	defer cancelCtx()

	done := make(chan output, 1)
	go func() {
		var out output = output{}
		switch in.Action {
		case services_protocols.ActionUserCreate:
			out = svc.create(ctx, in.Payload.User)
		default:
			svc.log.Metadata(map[string]interface{}{"input": in}).Error("[UserService.Exec] unrecognized command")
			panic(fmt.Sprintf("[UserService.Exec] unrecognized command %s", in.Action))
		}

		done <- out
	}()

	select {
	case <-ctx.Done():
		svc.log.Metadata(map[string]interface{}{"input": in}).Error("[UserService.Exec] context done")
		panic("[UserService.Exec] context done")
	case out := <-done:
		return out
	}
}

func (svc *UserSvc) create(ctx context.Context, u entities.User) (out output) {
	ud := svc.usrRepo.Exec(ctx, repositories_protocols.UserRepositoryInput{
		Action:  repositories_protocols.ActionUserFindByEmail,
		Payload: struct{ entities.User }{u},
	}).UserDto

	if ud != nil {
		svc.log.Metadata(map[string]interface{}{"user": u}).Warn("[UserService.create] duplicated user")
		out.Error = errors_adapters.NewInvalidBadRequestError("invalid payload")
		return
	}

	out.Result.UserDto = svc.usrRepo.Exec(ctx, repositories_protocols.UserRepositoryInput{
		Action:  repositories_protocols.ActionUserCreate,
		Payload: struct{ entities.User }{u},
	}).UserDto

	return
}
