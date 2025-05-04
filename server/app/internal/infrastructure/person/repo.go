package person_infra

import (
	"context"
	"fmt"
	"person-details-service/internal/domain/person"
	"person-details-service/pkg/postgres"
)

type PersonRepository struct {
	db postgres.PgClient
}

func NewPersonRepository(db postgres.PgClient) PersonRepository {
	return PersonRepository{db}
}

func (r PersonRepository) SavePerson(ctx context.Context, p person.Person) error {
	const op = "PersonRepository.SavePerson: %w"

	const query = `
		INSERT INTO persons(id, name, surname, patronymic, age, gender, nationality, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`

	dto := MapPersonToDTO(p)

	_, err := r.db.Exec(
		ctx,
		query,
		dto.ID,
		dto.Name,
		dto.Surname,
		dto.Patronymic,
		dto.Age,
		dto.Gender,
		dto.Nationality,
		dto.CreatedAt,
		dto.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return nil
}
