package utils

import "github.com/pkg/errors"

type ValidationError interface {
	Error() string
}

type validationError struct {
	Msg string
}

func WrapValidationError(err error, errorMessage string) ValidationError {
	return &validationError{Msg: errors.Wrap(err, errorMessage).Error()}
}

func (ve *validationError) Error() string {
	return ve.Msg
}

func IsValidationError(err error) bool {
	switch err.(type) {
	case *validationError:
		return true
	}
	return false
}