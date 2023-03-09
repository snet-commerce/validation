package validation

import (
	"encoding/json"
	"fmt"
)

// violation represents some validation rule violation
type violation struct {
	reason  string
	message string
}

// Violation creates new violation
func Violation(reason, msg string) violation {
	return violation{
		reason:  reason,
		message: msg,
	}
}

// String implements fmt.Stringer interface
func (v violation) String() string {
	return fmt.Sprintf("%s: %s", v.reason, v.message)
}

// MarshalJSON implements json.Marshaler interface
func (v violation) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Reason  string `json:"reason"`
		Message string `json:"message"`
	}{
		Reason:  v.reason,
		Message: v.message,
	})
}
