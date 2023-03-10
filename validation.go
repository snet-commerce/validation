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

// Err creates validation error or nil if no violations occurred
func (r *Result) Err() error {
	if !r.HasError() {
		return nil
	}

	return &Error{
		violations: r.violations,
	}
}
