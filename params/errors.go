package params

type NoCountryError struct {
	msg string
}

func (e NoCountryError) Error() string {
	return e.msg
}
