package errors_adapters

import errors_protocols "github.com/vm-championships-manager/backend-service/internal/errors/protocols"

const (
	EntityValidationErrorName = "EntityValidationError"
	EntityValidationErrorCode = 422
)

func NewEntityValidationError(msg string) errors_protocols.CustomError {
	return &CustomErrorAdapter{
		Name:    EntityValidationErrorName,
		Code:    EntityValidationErrorCode,
		Message: msg,
	}
}
