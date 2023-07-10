package utils_protocols

type Log interface {
	Metadata(map[string]interface{}) Log
	Error(msg string)
	Warn(msg string)
}
