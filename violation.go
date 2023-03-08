package validation

import (
	"encoding/json"
	"fmt"
)

type violation struct {
	reason  string
	message string
}

func Violation(reason, msg string) violation {
	return violation{
		reason:  reason,
		message: msg,
	}
}

func (v violation) Error() string {
	return fmt.Sprintf("%s: %s", v.reason, v.message)
}

func (v violation) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Reason  string `json:"reason"`
		Message string `json:"message"`
	}{
		Reason:  v.reason,
		Message: v.message,
	})
}
