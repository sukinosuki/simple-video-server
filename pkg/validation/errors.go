package validation

type ValidationError struct {
	Msg string
}

func (v *ValidationError) Error() string {

	return v.Msg
}

func NewValidateError(msg string) *ValidationError {

	return &ValidationError{
		msg,
	}
}
