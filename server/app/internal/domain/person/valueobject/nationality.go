package person_vo

import (
	failure "person-details-service/internal/domain/person/failure"
	"strings"
	"unicode/utf8"
)

type Nationality struct {
	value string
}

func NewNationality(nationality string) (*Nationality, error) {
	if utf8.RuneCountInString(nationality) != 2 {
		return nil, failure.ErrUnexpectedNationality
	}
	return &Nationality{strings.ToTitle(nationality)}, nil
}

func (n Nationality) Value() string {
	return n.value
}

func (n Nationality) Equals(other Nationality) bool {
	return n.Value() == other.Value()
}
