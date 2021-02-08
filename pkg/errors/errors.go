package errors

import "github.com/Zzocker/bookolab/pkg/code"

// E represents error used in this package
type E interface {
	// GetStatus return status of error
	GetStatus() code.Status

	// InError : internal error to be logged
	InError() error

	// Message to be sent to client
	Message() string
}

// Init : create new error
// we can say like initiating customs error
func Init(err error, status code.Status, msg string) E {
	return nil
}
