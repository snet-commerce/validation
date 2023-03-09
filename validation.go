package validation

// Result represents collection of violations
type Result struct {
	violations []violation
}

// NewResult builds new Result
func NewResult(violations ...violation) *Result {
	return &Result{violations: violations}
}

// HasError indicates if Result has any violations
func (r *Result) HasError() bool {
	return len(r.violations) > 0
}

// Err creates validation error with no message or nil if no violations occurred
func (r *Result) Err() error {
	return r.ErrWithMsg("")
}

// ErrWithMsg creates validation error with corresponding message or nil if no violations occurred
func (r *Result) ErrWithMsg(msg string) error {
	if !r.HasError() {
		return nil
	}

	return &Error{
		message:    msg,
		violations: r.violations,
	}
}
