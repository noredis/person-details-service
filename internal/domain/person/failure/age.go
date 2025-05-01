package person_failure

var ErrAgeIsNegative = errAgeIsNegative()

type AgeIsNegative struct{}

func (e AgeIsNegative) Error() string {
	return "age should not be less than 0"
}

func errAgeIsNegative() error {
	return AgeIsNegative{}
}
