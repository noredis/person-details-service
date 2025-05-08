package person_repo

import (
	"context"
	"fmt"
	"person-details-service/internal/domain/person"
	vo "person-details-service/internal/domain/person/valueobject"
	"slices"
)

type FilterOptions struct {
	Age         *vo.Age
	Gender      *vo.Gender
	Nationality *vo.Nationality
}

type PersonRepository interface {
	SavePerson(ctx context.Context, p person.Person) error
	GetPersonByID(ctx context.Context, id vo.PersonID) (*person.Person, error)
	UpdatePerson(ctx context.Context, p person.Person) error
	DeletePerson(ctx context.Context, id vo.PersonID) error
	GetPersons(ctx context.Context, filterOptions FilterOptions) ([]person.Person, error)
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
		return nil, fmt.Errorf("err")
	}

	p := r.persons[idx]
	return &p, nil
}

func (r *FakePersonRepository) UpdatePerson(ctx context.Context, p person.Person) error {
	idx := slices.IndexFunc(r.persons, func(pers person.Person) bool {
		return pers.ID().Equals(p.ID())
	})

	if idx < 0 {
		return nil
	}

	r.persons[idx] = p
	return nil
}

func (r *FakePersonRepository) DeletePerson(ctx context.Context, id vo.PersonID) error {
	idx := slices.IndexFunc(r.persons, func(p person.Person) bool {
		return p.ID().Equals(id)
	})

	if idx < 0 {
		return nil
	}

	r.persons = slices.Delete(r.persons, idx, idx+1)
	return nil
}

func (r *FakePersonRepository) GetPersons(ctx context.Context, filterOptions FilterOptions) ([]person.Person, error) {
	persons := make([]person.Person, 0)

	for _, p := range r.persons {
		if matchesFilters(p, filterOptions) {
			persons = append(persons, p)
		}
	}

	return persons, nil
}

func matchesFilters(p person.Person, filters FilterOptions) bool {
	if filters.Age != nil && p.Age() != nil && p.Age().Value() != filters.Age.Value() {
		return false
	}

	if filters.Gender != nil && p.Gender() != nil && p.Gender().Value() != filters.Gender.Value() {
		return false
	}

	if filters.Nationality != nil && p.Nationality() != nil && p.Nationality().Value() != filters.Nationality.Value() {
		return false
	}

	return true
}
