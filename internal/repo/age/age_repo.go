package age_repo

import (
	"context"
	"person-details-service/internal/domain/person"
	vo "person-details-service/internal/domain/person/valueobject"
)

type AgeRepository interface {
	FindOutPersonsAge(ctx context.Context, p person.Person) (*vo.Age, error)
}

type FakeAgeRepository struct{}

func (r FakeAgeRepository) FindOutPersonsAge(ctx context.Context, p person.Person) (*vo.Age, error) {
	if p.FullName() == "John Doe" {
		return vo.NewAge(37)
	}
	return nil, nil
}
