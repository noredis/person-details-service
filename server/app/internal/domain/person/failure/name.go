package person_failure

var ErrNameIsEmpty = errNameIsEmpty()

type NameIsEmpty struct{}

func (e NameIsEmpty) Error() string {
	return "name should not be empty"
}

func errNameIsEmpty() error {
	return NameIsEmpty{}
}
