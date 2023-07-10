package services_protocols

import (
	"context"

	"github.com/vm-championships-manager/backend-service/internal/dtos"
	"github.com/vm-championships-manager/backend-service/internal/entities"
	errors_protocols "github.com/vm-championships-manager/backend-service/internal/errors/protocols"
)

const (
	ActionUserCreate = "USER/CREATE"
)

type UserServiceInput struct {
	Action  string
	Payload struct{ entities.User }
}

type UserServiceOutput struct {
	Error  errors_protocols.CustomError
	Result struct{ *dtos.UserDto }
}

type UserService interface {
	Exec(ctx context.Context, in UserServiceInput) UserServiceOutput
}
