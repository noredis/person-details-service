package person

import (
	vo "person-details-service/internal/domain/person/valueobject"
	"time"
)

type Person struct {
	id          vo.PersonID
	name        vo.Name
	surname     vo.Name
	patronymic  *vo.Patronymic
	age         *vo.Age
	gender      *vo.Gender
	nationality *vo.Nationality
	createdAt   time.Time
	updatedAt   *time.Time
}

func CreatePerson(
	id vo.PersonID,
	name vo.Name,
	surname vo.Name,
	patronymic *vo.Patronymic,
	now time.Time,
) *Person {
	return &Person{
		id:          id,
		name:        name,
		surname:     surname,
		patronymic:  patronymic,
		age:         nil,
		gender:      nil,
		nationality: nil,
		createdAt:   now,
		updatedAt:   nil,
	}
}

func (p *Person) ID() vo.PersonID {
	return p.id
}

func (p *Person) Name() vo.Name {
	return p.name
}

func (p *Person) Surname() vo.Name {
	return p.surname
}

func (p *Person) Patronymic() *vo.Patronymic {
	return p.patronymic
}

func (p *Person) FullName() vo.FullName {
	return vo.NewFullName(p.Name(), p.Surname(), p.Patronymic())
}

func (p *Person) Age() *vo.Age {
	return p.age
}

func (p *Person) Gender() *vo.Gender {
	return p.gender
}

func (p *Person) Nationality() *vo.Nationality {
	return p.nationality
}

func (p *Person) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Person) UpdatedAt() *time.Time {
	return p.updatedAt
}

func (p *Person) SpecifyAge(age *vo.Age) {
	p.age = age
}

func (p *Person) SpecifyGender(gender *vo.Gender) {
	p.gender = gender
}

func (p *Person) SpecifyNationality(nationality *vo.Nationality) {
	p.nationality = nationality
}
