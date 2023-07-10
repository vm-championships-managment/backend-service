package pkg_protocols

type ErrorHandler interface {
	// Returns output function or log and panic if error is not nil
	Double(result interface{}, err error) func(errMsg string, metadataErr map[string]interface{}) interface{}
	// Log and panic if error is not nil
	Single(err error) func(errMsg string, metadataErr map[string]interface{})
}
