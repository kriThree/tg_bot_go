package utils

import "errors"

type MiddlewareError struct {
	e error
}

func (m MiddlewareError) Error() string {
	return m.e.Error()
}

func New(e string) MiddlewareError {
	return MiddlewareError{e: errors.New(e)}
}

var (
	NotValidUpdateErr = New("not valid update - not found user ID")
)
