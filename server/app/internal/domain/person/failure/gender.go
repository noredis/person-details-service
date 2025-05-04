package person_failure

var ErrUnexpectedGender = errUnexpectedGender()

type UnexpectedGender struct{}

func (e UnexpectedGender) Error() string {
	return "gender should be male or female"
}

func errUnexpectedGender() error {
	return UnexpectedGender{}
}
