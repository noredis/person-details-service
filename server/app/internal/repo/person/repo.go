package person_repo

import (
	"context"
	"person-details-service/internal/domain/person"
)

type PersonRepository interface {
	SavePerson(ctx context.Context, p person.Person) error
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
