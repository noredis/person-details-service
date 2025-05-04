package age_repo

import (
	"context"
	"person-details-service/internal/domain/person/valueobject"
)

type AgeRepository interface {
	FindOutPersonsAge(ctx context.Context, fullName person_vo.FullName) (*person_vo.Age, error)
}

type FakeAgeRepository struct{}

func (r FakeAgeRepository) FindOutPersonsAge(ctx context.Context, fullName person_vo.FullName) (*person_vo.Age, error) {
	if fullName.Value() == "John Doe" {
		return person_vo.NewAge(37)
	}
	return nil, nil
}
