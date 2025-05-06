package person_infra

import (
	"database/sql"
	"person-details-service/internal/domain/person"
	"time"
)

type PersonDTO struct {
	ID          string
	Name        string
	Surname     string
	Patronymic  sql.NullString
	Age         sql.NullInt16
	Gender      sql.NullString
	Nationality sql.NullString
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func MapPersonToDTO(p person.Person) PersonDTO {
	dto := PersonDTO{
		ID:        p.ID().Value(),
		Name:      p.Name().Value(),
		Surname:   p.Surname().Value(),
		CreatedAt: p.CreatedAt(),
		UpdatedAt: p.UpdatedAt(),
	}

	if p.Patronymic() != nil {
		dto.Patronymic.String = p.Patronymic().Value()
		dto.Patronymic.Valid = true
	}

	if p.Age() != nil {
		dto.Age.Int16 = int16(p.Age().Value())
		dto.Age.Valid = true
	}

	if p.Gender() != nil {
		dto.Gender.String = p.Gender().Value()
		dto.Gender.Valid = true
	}

	if p.Nationality() != nil {
		dto.Nationality.String = p.Nationality().Value()
		dto.Nationality.Valid = true
	}

	return dto
}
