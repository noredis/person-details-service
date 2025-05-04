package person_vo

import (
	"fmt"
	failure "person-details-service/internal/domain/person/failure"
	"strings"
)

type Name struct {
	value string
}

func NewName(name string) (*Name, error) {
	const op = "person_vo.NewName: %w"

	name = strings.Trim(name, " ")
	if name == "" {
		return nil, fmt.Errorf(op, failure.ErrNameIsEmpty)
	}

	return &Name{name}, nil
}

func (n Name) Value() string {
	return n.value
}

func (n Name) Equals(other Name) bool {
	return n.Value() == other.Value()
}
