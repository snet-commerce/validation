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

// violation represents validation business rule violation
type violation struct {
	message string
	code    string
	args    map[string]string
}

// Violation creates new violation
func Violation(msg, code string, args ...MessageArg) violation {
	m := make(map[string]string)
	for _, a := range args {
		m[a.Key] = a.Value
	}

	return violation{
		message: msg,
		code:    code,
		args:    m,
	}
}

// MarshalJSON implements json.Marshaler interface
func (v violation) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Message string            `json:"message"`
		Code    string            `json:"code"`
		Args    map[string]string `json:"args"`
	}{
		Message: v.message,
		Code:    v.code,
		Args:    v.args,
	})
}

// String implements fmt.Stringer interface
func (v violation) String() string {
	args := make([]string, len(v.args))
	for key, val := range v.args {
		args = append(args, fmt.Sprintf("%s = %s", key, val))
	}

	var s string
	if len(args) > 0 {
		s = fmt.Sprintf(" - args: %s", strings.Join(args, ", "))
	}

	return fmt.Sprintf("code %s - %s%s", v.code, v.message, s)
}
