package repositores_protocols

import (
	"context"

	"github.com/vm-championships-manager/backend-service/internal/dtos"
	"github.com/vm-championships-manager/backend-service/internal/entities"
)

const (
	ActionUserCreate      = "USER/CREATE"
	ActionUserFindByEmail = "USER/FIND_BY_EMAIL"
)

type UserRepositoryInput struct {
	Action  string
	Payload struct{ entities.User }
}

type UserRepositoryOutput struct {
	UserDto *dtos.UserDto
}

type UserRepository interface {
	Exec(ctx context.Context, in UserRepositoryInput) UserRepositoryOutput
}
