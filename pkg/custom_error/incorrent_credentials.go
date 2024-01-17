package custom_error

type IncorrectCredentialsError struct {
}

func (e *IncorrectCredentialsError) Error() string {
	return "incorrect phone_number or password"
}
