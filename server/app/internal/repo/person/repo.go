package person_repo

import (
	"context"
	"person-details-service/internal/domain/person"
	vo "person-details-service/internal/domain/person/valueobject"
	"slices"
)

type PersonRepository interface {
	SavePerson(ctx context.Context, p person.Person) error
	GetPersonByID(ctx context.Context, id vo.PersonID) (*person.Person, error)
}

type FakePersonRepository struct {
	persons []person.Person
}

func NewFakePersonRepository() *FakePersonRepository {
	persons := make([]person.Person, 0)
	return &FakePersonRepository{persons}
}

func (r *FakePersonRepository) SavePerson(ctx context.Context, p person.Person) error {
	r.persons = append(r.persons, p)
	return nil
}

func (r *FakePersonRepository) GetPersonByID(ctx context.Context, id vo.PersonID) (*person.Person, error) {
	idx := slices.IndexFunc(r.persons, func(p person.Person) bool {
		return p.ID().Equals(id)
	})

	if idx < 0 {
		return nil, nil
	}

	p := r.persons[idx]
	return &p, nil
}
