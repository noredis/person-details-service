package person_vo

import (
	failure "person-details-service/internal/domain/person/failure"
	"strings"
)

type Surname struct {
	value string
}

func NewSurname(surname string) (*Surname, error) {
	surname = strings.Trim(surname, " ")
	if surname == "" {
		return nil, failure.ErrSurnameIsEmpty
	}

	return &Surname{surname}, nil
}

func (s Surname) Value() string {
	return s.value
}

func (s Surname) Equals(other Surname) bool {
	return s.Value() == other.Value()
}
