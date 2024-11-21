package errors

import (
	"errors"
	"fmt"

	"go.tradeforge.dev/fmp/model"
)

// ResponseError represents an API response with an error status code.
type ResponseError struct {
	model.BaseResponse

	// An HTTP status code for unsuccessful requests.
	StatusCode int
}

// Error returns the details of an error response.
func (e *ResponseError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("request %s failed with code %d: %s", e.RequestID, e.StatusCode, e.Message)
	}
	return fmt.Sprintf("request %s failed with code %d", e.RequestID, e.StatusCode)
}

func AsResponseError(obj any) (*ResponseError, bool) {
	err, ok := obj.(error)
	if !ok {
		return nil, false
	}

	responseError := &ResponseError{}
	if ok := errors.As(err, &responseError); !ok {
		return nil, false
	}
	return responseError, true
}

// Code is useful for converting to HTTP status code.
// In general it's a value which is machine readable.
type Code string

// Error represents API error. All fields (except of wrapped error and message)
// are meant to be publicly sharable.
type Error struct {
	Err           error
	Message       string
	PublicMessage string
	Code          Code
	Data          any
}

func NewError(message string, code Code) *Error {
	return &Error{
		Err:           nil,
		Message:       message,
		PublicMessage: "",
		Code:          code,
		Data:          nil,
	}
}

func NewPublicError(publicMessage string, code Code) *Error {
	return &Error{
		Err:           nil,
		Message:       publicMessage,
		PublicMessage: publicMessage,
		Code:          code,
		Data:          nil,
	}
}

func (e *Error) Error() string {
	if e.Message != "" {
		if e.Err != nil {
			return fmt.Sprintf("%s: %s", e.Message, e.Err.Error())
		}
		return e.Message
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return ""
}

func (e *Error) Wrap(err error) *Error {
	cloned := e.clone()
	cloned.Err = err
	return cloned
}

func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) WithPublicMessage(msg string) *Error {
	cloned := e.clone()
	cloned.PublicMessage = msg
	return cloned
}

func (e *Error) WithData(data any) *Error {
	cloned := e.clone()
	cloned.Data = data
	return cloned
}

func (e *Error) clone() *Error {
	cloned := *e
	return &cloned
}

func IsErrorWithCode(err error, code Code) bool {
	var apiErr *Error
	if ok := errors.As(err, &apiErr); ok {
		return apiErr.Code == code
	}
	return false
}
