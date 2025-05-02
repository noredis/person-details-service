package person

import (
	vo "person-details-service/internal/domain/person/valueobject"
)

type Person struct {
	name       vo.Name
	surname    vo.Name
	patronymic *vo.Patronymic
	age        *vo.Age
}

func CreatePerson(name vo.Name, surname vo.Name, patronymic *vo.Patronymic) *Person {
	return &Person{
		name:       name,
		surname:    surname,
		patronymic: patronymic,
		age:        nil,
	}
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

func (p *Person) SpecifyAge(age *vo.Age) {
	p.age = age
}
