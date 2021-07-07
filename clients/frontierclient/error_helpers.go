package frontierclient

import "github.com/xdbfoundation/go/support/errors"

// IsNotFoundError returns true if the error is a frontierclient.Error with
// a not_found problem indicating that the resource is not found on
// Frontier.
func IsNotFoundError(err error) bool {
	var hErr *Error

	err = errors.Cause(err)
	switch err := err.(type) {
	case *Error:
		hErr = err
	case Error:
		hErr = &err
	}

	if hErr == nil {
		return false
	}

	return hErr.Problem.Type == "https://digitalbits.org/frontier-errors/not_found"
}

// GetError returns an error that can be interpreted as a frontier-specific
// error. If err cannot be interpreted as a frontier-specific error, a nil error
// is returned. The caller should still check whether err is nil.
func GetError(err error) *Error {
	var hErr *Error

	err = errors.Cause(err)
	switch e := err.(type) {
	case *Error:
		hErr = e
	case Error:
		hErr = &e
	}

	return hErr
}
