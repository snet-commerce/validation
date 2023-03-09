package validation

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Error represents error produced by validation result
type Error struct {
	message    string
	violations []violation
}

// Error implements builtin error interface
func (e *Error) Error() string {
	var prelude string
	if e.message != "" {
		prelude = fmt.Sprintf("%s: ", prelude)
	}

	errors := make([]string, len(e.violations))
	for _, v := range e.violations {
		errors = append(errors, v.String())
	}

	return fmt.Sprintf("%s%s", prelude, strings.Join(errors, "; "))
}

// MarshalJSON implements json.Marshaler interface
func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Message string      `json:"message,omitempty"`
		Errors  []violation `json:"errors"`
	}{
		Message: e.message,
		Errors:  e.violations,
	})
}
