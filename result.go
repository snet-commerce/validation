package validation

// Result represents collection of errors
type Result struct {
	errors []error
}

// NewResult builds new Result
func NewResult() *Result {
	return &Result{errors: make([]error, 0)}
}

// HasError indicates if Result has any violations
func (r *Result) HasError() bool {
	return len(r.errors) > 0
}

// Error adds new error
func (r *Result) Error(err error) *Result {
	if err != nil {
		r.errors = append(r.errors, err)
	}
	return r
}

// RaiseErr creates ResultError or nil if no errors occurred on validation
func (r *Result) RaiseErr() error {
	if !r.HasError() {
		return nil
	}
	return &ResultError{
		errors: r.errors,
	}
}
