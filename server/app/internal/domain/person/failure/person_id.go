package person_failure

var ErrParsePersonID = errParsePersonID()

type ParsePersonID struct{}

func (e ParsePersonID) Error() string {
	return "unable to parse person id"
}

func errParsePersonID() error {
	return ParsePersonID{}
}
