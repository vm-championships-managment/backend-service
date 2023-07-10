package errors_protocols

type CustomErrorInfo struct {
	Name    string
	Message string
	Code    uint16
}

type CustomError interface {
	GetErrorInfos() CustomErrorInfo
}
