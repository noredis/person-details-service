package age_repo

import (
	"context"
	"fmt"
	"person-details-service/internal/domain/person/valueobject"
)

type AgeRepository interface {
	FindOutPersonsAge(ctx context.Context, fullName person_vo.FullName) (*person_vo.Age, error)
}

type FakeAgeRepository struct{}

func (r FakeAgeRepository) FindOutPersonsAge(ctx context.Context, fullName person_vo.FullName) (*person_vo.Age, error) {
	if fullName.Value() == "John Doe" {
		return person_vo.NewAge(37)
	} else if fullName.Value() == "John Doe 1" {
		return nil, fmt.Errorf("err")
	}
	return nil, nil
}
