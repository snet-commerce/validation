package validation

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MessageArg represents message placeholder key and its value
type MessageArg struct {
	Key   string
	Value string
}

// Error represents validation violation
type Error struct {
	message string
	code    string
	args    map[string]string
}

// NewError constructs new error
func NewError(msg, code string, args ...MessageArg) error {
	m := make(map[string]string)
	for _, a := range args {
		m[a.Key] = a.Value
	}
	return &Error{
		code:    code,
		message: msg,
		args:    m,
	}
}

// Error implements builtin error interface
func (e *Error) Error() string {
	args := make([]string, 0, len(e.args))
	for key, val := range e.args {
		args = append(args, fmt.Sprintf("%s = %s", key, val))
	}

	var s string
	if len(args) > 0 {
		s = fmt.Sprintf(" - args: %s", strings.Join(args, ", "))
	}

	return fmt.Sprintf("code %s - %s%s", e.code, e.message, s)
}

// MarshalJSON implements json.Marshaler interface
func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Message string            `json:"message"`
		Code    string            `json:"code"`
		Args    map[string]string `json:"args"`
	}{
		Message: e.message,
		Code:    e.code,
		Args:    e.args,
	})
}

// ResultError represents validation result error
type ResultError struct {
	errors []error
}

// Error implements builtin error interface
func (r *ResultError) Error() string {
	var sb strings.Builder
	for i, e := range r.errors {
		if i != 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(fmt.Sprintf("%s;", e))
	}
	return sb.String()
}

// MarshalJSON implements json.Marshaler interface
func (r *ResultError) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Errors []error `json:"errors"`
	}{
		Errors: r.errors,
	})
}
