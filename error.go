package validation

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Error represents error produced by validation result
type Error struct {
	violations []violation
}

// Error implements builtin error interface
func (e *Error) Error() string {
	var sb strings.Builder
	for _, v := range e.violations {
		sb.WriteString(fmt.Sprintf("%s;", v.String()))
	}
	return sb.String()
}

// MarshalJSON implements json.Marshaler interface
func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Errors []violation `json:"errors"`
	}{
		Errors: e.violations,
	})
}
