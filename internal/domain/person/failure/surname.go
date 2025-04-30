package person_failure

var ErrSurnameIsEmpty = errSurnameIsEmpty()

type SurnameIsEmpty struct{}

func (e SurnameIsEmpty) Error() string {
	return "surname should not be empty"
}

func errSurnameIsEmpty() error {
	return SurnameIsEmpty{}
}
