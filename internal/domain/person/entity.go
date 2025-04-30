package person

import (
	"fmt"
	vo "person-details-service/internal/domain/person/valueobject"
)

type Person struct {
	name       vo.Name
	surname    vo.Name
	patronymic *vo.Patronymic
}

func CreatePerson(name vo.Name, surname vo.Name, patronymic *vo.Patronymic) *Person {
	return &Person{name, surname, patronymic}
}

func (p Person) Name() vo.Name {
	return p.name
}

func (p Person) Surname() vo.Name {
	return p.surname
}

func (p Person) Patronymic() *vo.Patronymic {
	return p.patronymic
}

func (p Person) FullName() string {
	if p.Patronymic() == nil {
		return fmt.Sprintf("%s %s", p.Name().Value(), p.Surname().Value())
	}
	return fmt.Sprintf("%s %s %s", p.Name().Value(), p.Surname().Value(), p.Patronymic().Value())
}
