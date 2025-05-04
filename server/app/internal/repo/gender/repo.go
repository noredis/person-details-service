package gender_repo

import (
	"context"
	"person-details-service/internal/domain/person/valueobject"
)

type GenderRepository interface {
	FindOutPersonsGender(ctx context.Context, fullName person_vo.FullName) (*person_vo.Gender, error)
}

type FakeGenderRepository struct{}

func (r FakeGenderRepository) FindOutPersonsGender(ctx context.Context, fullName person_vo.FullName) (*person_vo.Gender, error) {
	if fullName.Value() == "John Doe" {
		return person_vo.NewGender("male")
	}
	return nil, nil
}
