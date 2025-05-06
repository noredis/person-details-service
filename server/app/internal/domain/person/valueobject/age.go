package person_vo

import (
	"fmt"
	failure "person-details-service/internal/domain/person/failure"
)

type Age struct {
	value int
}

func NewAge(age int) (*Age, error) {
	const op = "person_vo.NewAge: %w"

	if age < 0 {
		return nil, fmt.Errorf(op, failure.ErrAgeIsNegative)
	}
	return &Age{age}, nil
}

func (a Age) Value() int {
	return a.value
}

func (a Age) Equals(other Age) bool {
	return a.Value() == other.Value()
}
