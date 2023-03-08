package validation

type Result struct {
	violations []violation
}

func NewResult(violations ...violation) *Result {
	return &Result{violations: violations}
}

func (r *Result) Err(msg string) error {
	if len(r.violations) == 0 {
		return nil
	}

	return &Error{
		message:    msg,
		violations: r.violations,
	}
}
