package errs

import (
	"fmt"
)

type Error struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"error"` // 系统内置的错误，不能使用该再短
	Info      string `json:"info"`  // 用于dev调试，但这个后续看线上要用。例如一个参数错误，可能有多种原因
	Status    int    `json:"-"`
	Values    []any  `json:"-"`
	Err       error  `json:"-"`
}

func NewError(status, code int) *Error {
	return &Error{
		Status:    status,
		ErrorCode: code,
	}
}

func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	switch err.(type) {
	case *Error:
		return err.(*Error)
	default:
		errUnknown := ErrUnknown.WithRawError(err)
		errUnknown.Message = err.Error()
		return errUnknown
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("code=%d. message=%s", e.ErrorCode, e.Message)
}

// WithInfo 用于Debug信息
func (e *Error) WithInfo(messageFormat string, a ...any) *Error {
	err := *e
	err.Info = fmt.Sprintf(messageFormat, a...)
	return &err
}

func (e *Error) WithRawError(err error) *Error {
	aerr := *e
	aerr.Err = err
	return &aerr
}

func (e *Error) WithValues(values ...any) *Error {
	err := *e
	err.Values = values
	return &err
}

func (e *Error) Clone() *Error {
	res := *e
	return &res
}

func (e *Error) WithMessage(message string) *Error {
	err := *e
	err.Message = message
	return &err
}
