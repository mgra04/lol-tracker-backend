package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
    Code    int
    Message string
    Err     error
}

func (e *AppError) Error() string {
    if e.Err != nil {
        return fmt.Sprintf("%s: %v", e.Message, e.Err)
    }
    return e.Message
}

func New(msg string, code int) *AppError {
    return &AppError{Message: msg, Code: code}
}

func Wrap(err error, msg string, code int) *AppError {
    if err == nil {
        return nil
    }
    return &AppError{Message: msg, Code: code, Err: err}
}

func IsNotFound(err error) bool {
    if appErr, ok := err.(*AppError); ok {
        return appErr.Code == http.StatusNotFound
    }
    return false
}
