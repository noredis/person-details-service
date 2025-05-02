package nationality_repo

import (
	"context"
	"person-details-service/internal/domain/person/valueobject"
)

type NationalityRepository interface {
	FindOutPersonsNationality(ctx context.Context, fullName person_vo.FullName) (*person_vo.Nationality, error)
}

type FakeNationalityRepository struct{}

func (r FakeNationalityRepository) FindOutPersonsNationality(ctx context.Context, fullName person_vo.FullName) (*person_vo.Nationality, error) {
	if fullName.Value() == "John Doe" {
		return person_vo.NewNationality("EN")
	}
	return nil, nil
}
