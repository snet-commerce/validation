package validation

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Error struct {
	message    string
	violations []violation
}

func (e *Error) Error() string {
	errors := make([]string, len(e.violations))
	for _, v := range e.violations {
		errors = append(errors, v.Error())
	}
	return fmt.Sprintf("%s: %s", e.message, strings.Join(errors, "; "))
}

func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Message string      `json:"message"`
		Errors  []violation `json:"errors"`
	}{
		Message: e.message,
		Errors:  e.violations,
	})
}
