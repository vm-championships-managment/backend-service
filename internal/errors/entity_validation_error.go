package internal_errors

import "fmt"

func EntityValidationError(message string) error {
	return fmt.Errorf("[internal_error][entity_validation_error] - some fields are invalids: %s", message)
}
