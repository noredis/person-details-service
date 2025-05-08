package nationality_repo

import (
	"context"
	"fmt"
	"person-details-service/internal/domain/person/valueobject"
)

type NationalityRepository interface {
	FindOutPersonsNationality(ctx context.Context, fullName person_vo.FullName) (*person_vo.Nationality, error)
}

type FakeNationalityRepository struct{}

func (r FakeNationalityRepository) FindOutPersonsNationality(ctx context.Context, fullName person_vo.FullName) (*person_vo.Nationality, error) {
	if fullName.Value() == "John Doe" {
		return person_vo.NewNationality("EN")
	} else if fullName.Value() == "John Doe 3" {
		return nil, fmt.Errorf("err")
	}
	return nil, nil
}
