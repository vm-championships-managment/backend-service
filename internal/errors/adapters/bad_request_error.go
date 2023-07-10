package errors_adapters

import errors_protocols "github.com/vm-championships-manager/backend-service/internal/errors/protocols"

const (
	InvalidPayloadErrorName = "BadRequest"
	InvalidPayloadErrorCode = 400
)

func NewInvalidBadRequestError(msg string) errors_protocols.CustomError {
	return &CustomErrorAdapter{
		Name:    InvalidPayloadErrorName,
		Code:    InvalidPayloadErrorCode,
		Message: msg,
	}
}
