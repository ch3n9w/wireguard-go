package device

import (
	"errors"
)

type errorTransformation struct {
	Message string `default:"this is not for error"`
	data    []string
}

func (e *errorTransformation) Error() string {
	return e.Message
}

func GetErrorTransformation(s []string) error {
	return &errorTransformation{data: s}
}

func CheckIsErrorTransformation(err error) bool {
	var errorTransformation *errorTransformation
	ok := errors.As(err, &errorTransformation)
	return ok
}

func GetData(err error) []string {
	var et *errorTransformation
	if errors.As(err, &et) {
		return et.data
	}
	return []string{}
}
