package domain

import (
	"fmt"
)

const (
	ErrCONFLICT       = "conflict"
	ErrINTERNAL       = "internal"
	ErrINVALID        = "invalid"
	ErrNOTFOUND       = "not_found"
	ErrNOTIMPLEMENTED = "not_implemented"
	ErrUNAUTHORIZED   = "unauthorized"
)

type Error struct {
	Code    string
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("wtf error: code=%s message=%s", e.Code, e.Message)
}

func Errorf(code string, format string, args ...interface{}) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}
