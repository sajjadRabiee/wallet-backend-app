package custom_error

type NotValidPhoneNumberError struct {
}

func (e *NotValidPhoneNumberError) Error() string {
	return "not a valid phone_number"
}
