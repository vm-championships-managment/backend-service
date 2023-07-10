package errors_adapters

import (
	errors_protocols "github.com/vm-championships-manager/backend-service/internal/errors/protocols"
)

type CustomErrorAdapter struct {
	Name    string
	Message string
	Code    uint16
}

func (ce *CustomErrorAdapter) GetErrorInfos() errors_protocols.CustomErrorInfo {
	return errors_protocols.CustomErrorInfo{
		Name:    ce.Name,
		Message: ce.Message,
		Code:    ce.Code,
	}
}
