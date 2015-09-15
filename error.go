package smsintel

import (
	"fmt"
)

type Error interface {
	error

	Code() int64
	Message() string
	PreviousError() error
}

type ApiError struct {
	code          int64
	message       string
	previousError error
}

func (e ApiError) Error() string {
	return fmt.Sprintf("%d: %s", e.code, e.message)
}

func (e ApiError) Code() int64 {
	return e.code
}

func (e ApiError) Message() string {
	return e.message
}

func (e ApiError) PreviousError() error {
	return e.previousError
}

func NewApiError(code int64, message string, previousError error) Error {
	if e, ok := previousError.(Error); ok && e != nil {
		return e
	}

	return &ApiError{
		code:          code,
		message:       message,
		previousError: previousError,
	}
}
