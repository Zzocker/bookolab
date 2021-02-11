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

	Error() string
}

// Init : create new error
// we can say like initiating customs error
func Init(e error, status code.Status, msg string) E {
	// Implement me
	return &err{
		status: status,
		err:    e,
		msg:    msg,
	}
}

type err struct {
	status code.Status
	err    error
	msg    string
}

func (e *err) GetStatus() code.Status {
	return e.status
}
func (e *err) InError() error {
	return e.err
}
func (e *err) Message() string {
	return e.msg
}

func (e *err) Error() string {
	return e.err.Error()
}
