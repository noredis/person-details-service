package person_vo

import (
	failure "person-details-service/internal/domain/person/failure"
	"strings"
)

type Name struct {
	value string
}

func NewName(name string) (*Name, error) {
	name = strings.Trim(name, " ")
	if name == "" {
		return nil, failure.ErrNameIsEmpty
	}

	return &Name{name}, nil
}

func (n Name) Value() string {
	return n.value
}

func (n Name) Equals(other Name) bool {
	return n.Value() == other.Value()
}
