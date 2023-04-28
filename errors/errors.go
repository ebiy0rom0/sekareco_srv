package errors

import (
	"sekareco_srv/code"

	"github.com/ebiy0rom0/errors"
)

type appError struct {
	code code.ErrorCode
	err  error
}

// New returns error interface.
func New(code code.ErrorCode, err error) error {
	return &appError{
		code: code,
		err:  errors.WithStack(err),
	}
}

// Is the wrapper for the import package's errors.Wrap().
// It's returns same result.
func Wrap(err error, msg string) error {
	return errors.Wrap(err, msg)
}

// Is the wrapper for the import package's errors.Wrapf().
// It's returns same result.
func Wrapf(err error, format string, args ...any) error {
	return errors.Wrapf(err, format, args...)
}

// Is the wrapper for the import package's errors.Is().
// It's returns same result.
func Is(err, target error) bool { return errors.Is(err, target) }

// As the wrapper for the import package's errors.As().
// It's returns same result.
func As(err error, target any) bool { return errors.As(err, target) }

// Error returns error message.
// It's an implementation of the error interface.
func (e *appError) Error() string {
	return e.err.Error()
}

//
func (e *appError) Is(err error) bool {
	return errors.Is(e.code, err)
}

func (e *appError) Unwrap() error {
	return e.err
}

var _ error = (*appError)(nil)
