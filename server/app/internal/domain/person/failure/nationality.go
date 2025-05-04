package person_failure

var ErrUnexpectedNationality = errUnexpectedNationality()

type UnexpectedNationality struct{}

func (e UnexpectedNationality) Error() string {
	return "unexpected nationality"
}

func errUnexpectedNationality() error {
	return UnexpectedNationality{}
}
