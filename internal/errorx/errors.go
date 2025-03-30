package errorx

import (
	"fmt"
)

type AppError struct {
	StatusCode int
	ErrorCode  string
	Message    string
}

func (e *AppError) Error() string {
	return e.Message
}

func New(statusCode int, errorCode, message string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		ErrorCode:  errorCode,
		Message:    message,
	}
}

func Wrap(err error, statusCode int, errorCode, message string) *AppError {
	if err == nil {
		return nil
	}

	return &AppError{
		StatusCode: statusCode,
		ErrorCode:  errorCode,
		Message:    fmt.Sprintf("%s: %v", message, err),
	}
}
