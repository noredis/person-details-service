package person_dto

import (
	"person-details-service/internal/domain/person"
	"time"
)

type PersonDTO struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Surname     string     `json:"surname"`
	Patronymic  *string    `json:"patronymic,omitempty"`
	Age         *int       `json:"age,omitempty"`
	Gender      *string    `json:"gender,omitempty"`
	Nationality *string    `json:"nationality,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func MapFromPerson(p person.Person) *PersonDTO {
	pDTO := PersonDTO{
		ID:        p.ID().Value(),
		Name:      p.Name().Value(),
		Surname:   p.Surname().Value(),
		CreatedAt: p.CreatedAt(),
		UpdatedAt: p.UpdatedAt(),
	}

	if p.Patronymic() != nil {
		patronymic := p.Patronymic().Value()
		pDTO.Patronymic = &patronymic
	}

	if p.Age() != nil {
		age := p.Age().Value()
		pDTO.Age = &age
	}

	if p.Gender() != nil {
		gender := p.Gender().Value()
		pDTO.Gender = &gender
	}

	if p.Nationality() != nil {
		nationality := p.Nationality().Value()
		pDTO.Nationality = &nationality
	}

	return &pDTO
}
