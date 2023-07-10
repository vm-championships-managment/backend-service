package errors_adapters

import errors_protocols "github.com/vm-championships-manager/backend-service/internal/errors/protocols"

const (
	InternalServerErrorName = "InternalServerError"
	InternalServerErrorCode = 500
)

func NewInternalServerError() errors_protocols.CustomError {
	return &CustomErrorAdapter{
		Name:    InternalServerErrorName,
		Code:    InternalServerErrorCode,
		Message: "internal server error",
	}
}
