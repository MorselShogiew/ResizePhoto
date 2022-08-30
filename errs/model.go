package errs

import (
	"encoding/json"
	"strings"
)

type AppError struct {
	Err *Err `json:"error"`
}

type Err struct {
	StatusCode      int    `json:"-"`
	Code            int    `json:"code"`
	Message         string `json:"message"`
	IsBusinessError bool   `json:"isBusinessError"`
	Error           string `json:"error"`
}

// New returns new AppError object
//
// `err` is real error caught by you;
// `errMsg` is your error desc;
// `isBusinessError` is flag of business error;
//
// If you have not an `err`, you can pass `nil` instead.
func New(err error, errMsg string, isBusinessError bool, statusCode int) *AppError {
	if err, ok := err.(*AppError); ok {
		return err
	}

	var errStr string
	if err != nil {
		errStr = err.Error()
	}

	return &AppError{
		Err: &Err{
			Message:         strings.ReplaceAll(errMsg, `"`, `\"`),
			IsBusinessError: isBusinessError,
			Error:           errStr,
			StatusCode:      statusCode,
		},
	}
}

// Error returns string representation of AppError in json format.
func (e *AppError) Error() string {
	out, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(out)
}

func (e *AppError) StatusCode() int {
	return e.Err.StatusCode
}
