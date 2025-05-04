package person_vo

import (
	failure "person-details-service/internal/domain/person/failure"
)

type Gender struct {
	value string
}

func NewGender(gender string) (*Gender, error) {
	switch gender {
	case "male", "female":
		return &Gender{gender}, nil
	}
	return nil, failure.ErrUnexpectedGender
}

func (g Gender) Value() string {
	return g.value
}

func (g Gender) IsMale() bool {
	return g.value == "male"
}

func (g Gender) IsFemale() bool {
	return g.value == "female"
}

func (g Gender) Equals(other Gender) bool {
	return g.Value() == other.Value()
}
