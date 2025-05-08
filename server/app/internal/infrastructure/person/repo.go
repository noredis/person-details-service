package person_infra

import (
	"context"
	"fmt"
	"person-details-service/internal/domain/person"
	vo "person-details-service/internal/domain/person/valueobject"
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

func (r PersonRepository) GetPersonByID(ctx context.Context, id vo.PersonID) (*person.Person, error) {
	const op = "PersonRepository.GetPersonByID: %w"

	const query = `
        SELECT id, name, surname, patronymic, age, gender, nationality, created_at, updated_at
        FROM persons
        WHERE id = $1;
    `

	var p PersonDTO

	row := r.db.QueryRow(ctx, query, id.Value())

	err := row.Scan(&p.ID, &p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Gender, &p.Nationality, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	name, err := vo.NewName(p.Name)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	surname, err := vo.NewName(p.Surname)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	var patronymic *vo.Patronymic
	if p.Patronymic.Valid {
		patronymic = vo.NewPatronymic(p.Patronymic.String)
	}

	var age *vo.Age
	if p.Age.Valid {
		age, err = vo.NewAge(int(p.Age.Int16))
		if err != nil {
			return nil, fmt.Errorf(op, err)
		}
	}

	var gender *vo.Gender
	if p.Gender.Valid {
		gender, err = vo.NewGender(p.Gender.String)
		if err != nil {
			return nil, fmt.Errorf(op, err)
		}
	}

	var nationality *vo.Nationality
	if p.Nationality.Valid {
		nationality, err = vo.NewNationality(p.Nationality.String)
		if err != nil {
			return nil, fmt.Errorf(op, err)
		}
	}

	restoredPerson := person.RestorePerson(id, *name, *surname, patronymic, age, gender, nationality, p.CreatedAt, p.UpdatedAt)

	return restoredPerson, nil
}

func (r PersonRepository) UpdatePerson(ctx context.Context, p person.Person) error {
	const op = "PersonRepository.UpdatePerson: %w"

	const query = `
		UPDATE persons SET
			name = $1,
			surname = $2,
			patronymic = $3,
			age = $4,
			gender = $5,
			nationality = $6,
			updated_at = $7
		WHERE id = $8;
	`

	dto := MapPersonToDTO(p)

	_, err := r.db.Exec(
		ctx,
		query,
		dto.Name,
		dto.Surname,
		dto.Patronymic,
		dto.Age,
		dto.Gender,
		dto.Nationality,
		dto.UpdatedAt,
		dto.ID,
	)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return nil
}

func (r PersonRepository) DeletePerson(ctx context.Context, id vo.PersonID) error {
	const op = "PersonRepository.DeletePerson: %w"

	const query = `
		DELETE FROM persons
		WHERE id = $1;
	`

	_, err := r.db.Exec(
		ctx,
		query,
		id.Value(),
	)
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return nil
}
