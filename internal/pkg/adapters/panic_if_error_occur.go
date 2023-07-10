package pkg_adapter

import (
	pkg_protocols "github.com/vm-championships-manager/backend-service/internal/pkg/protocols"
	utils_protocols "github.com/vm-championships-manager/backend-service/internal/utils/protocols"
)

type PanicIfErrorOccur struct {
	log utils_protocols.Log
}

func NewPanicIfErrorOccur(log utils_protocols.Log) pkg_protocols.ErrorHandler {
	return &PanicIfErrorOccur{
		log: log,
	}
}

// Returns output function or log and panic if error is not nil
func (piec *PanicIfErrorOccur) Double(result interface{}, err error) func(errMsg string, metadataErr map[string]interface{}) interface{} {
	return func(errMsg string, metadataErr map[string]interface{}) interface{} {
		if err != nil {
			piec.log.Metadata(map[string]interface{}{"metadata": metadataErr, "error": err}).Error(errMsg)
			panic(err)
		}
		return result
	}

}

// Panic if error is not nil
func (piec *PanicIfErrorOccur) Single(err error) func(errMsg string, metadata map[string]interface{}) {
	return func(errMsg string, metadata map[string]interface{}) {
		if err != nil {
			piec.log.Metadata(map[string]interface{}{"metadata": metadata, "error": err}).Error(errMsg)
			panic(err)
		}
	}
}
