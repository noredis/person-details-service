package person_vo

import (
	"fmt"
	failure "person-details-service/internal/domain/person/failure"

	"github.com/google/uuid"
)

type PersonID struct {
	value string
}

func NewPersonID() PersonID {
	return PersonID{uuid.New().String()}
}

func ParsePersonID(id string) (*PersonID, error) {
	const op = "person_vo.ParsePersonID: %w"

	_, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf(op, failure.ErrParsePersonID)
	}

	return &PersonID{id}, nil
}

func (pID PersonID) Value() string {
	return pID.value
}

func (pID PersonID) Equals(other PersonID) bool {
	return pID.Value() == other.Value()
}
