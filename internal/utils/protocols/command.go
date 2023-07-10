package utils_protocols

import errors_protocols "github.com/vm-championships-manager/backend-service/internal/errors/protocols"

type CommandGenericInput[T any] struct {
	Action  string
	Payload T
}

type CommandGenericOutput[T any] struct {
	Result T
	Error  errors_protocols.CustomError
}
