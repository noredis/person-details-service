package person_service

import (
	"context"
	"fmt"
	"person-details-service/internal/domain/person"
	vo "person-details-service/internal/domain/person/valueobject"
	"person-details-service/internal/repo/age"
	"person-details-service/internal/repo/gender"
	"person-details-service/internal/repo/nationality"
	"person-details-service/internal/repo/person"
	dto "person-details-service/internal/service/person/dto"
	"time"
)

type PersonService struct {
	ageRepo         age_repo.AgeRepository
	genderRepo      gender_repo.GenderRepository
	nationalityRepo nationality_repo.NationalityRepository
	personRepo      person_repo.PersonRepository
}

func NewPersonService(
	ageRepo age_repo.AgeRepository,
	genderRepo gender_repo.GenderRepository,
	nationalityRepo nationality_repo.NationalityRepository,
	personRepo person_repo.PersonRepository,
) PersonService {
	return PersonService{
		ageRepo:         ageRepo,
		genderRepo:      genderRepo,
		nationalityRepo: nationalityRepo,
		personRepo:      personRepo,
	}
}

func (s PersonService) CreatePerson(ctx context.Context, createPersonDTO dto.CreatePersonDTO) (*dto.PersonDTO, error) {
	const op = "PersonService.CreatePerson: %w"

	id := vo.NewPersonID()

	name, err := vo.NewName(createPersonDTO.Name)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	surname, err := vo.NewName(createPersonDTO.Surname)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	patronymic := vo.NewPatronymic(createPersonDTO.Patronymic)

	now := time.Now()

	p := person.CreatePerson(id, *name, *surname, patronymic, now)

	age, err := s.ageRepo.FindOutPersonsAge(ctx, p.FullName())
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	gender, err := s.genderRepo.FindOutPersonsGender(ctx, p.FullName())
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	nationality, err := s.nationalityRepo.FindOutPersonsNationality(ctx, p.FullName())
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	p.SpecifyAge(age)
	p.SpecifyGender(gender)
	p.SpecifyNationality(nationality)

	err = s.personRepo.SavePerson(ctx, *p)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	return dto.MapFromPerson(*p), nil
}

func (s PersonService) UpdatePerson(ctx context.Context, id string, updatePersonDTO dto.UpdatePersonDTO) (*dto.PersonDTO, error) {
	const op = "PersonService.UpdatePerson: %w"

	pID, err := vo.ParsePersonID(id)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	p, err := s.personRepo.GetPersonByID(ctx, *pID)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	name, err := vo.NewName(updatePersonDTO.Name)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	surname, err := vo.NewName(updatePersonDTO.Surname)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	patronymic := vo.NewPatronymic(updatePersonDTO.Patronymic)

	age, err := vo.NewAge(updatePersonDTO.Age)
	if err != nil && updatePersonDTO.Age != 0 {
		return nil, fmt.Errorf(op, err)
	}

	gender, err := vo.NewGender(updatePersonDTO.Gender)
	if err != nil && updatePersonDTO.Gender != "" {
		return nil, fmt.Errorf(op, err)
	}

	nationality, err := vo.NewNationality(updatePersonDTO.Nationality)
	if err != nil && updatePersonDTO.Nationality != "" {
		return nil, fmt.Errorf(op, err)
	}

	now := time.Now()

	p.EditPersonalInformation(*name, *surname, patronymic, age, gender, nationality, now)

	err = s.personRepo.UpdatePerson(ctx, *p)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	return dto.MapFromPerson(*p), nil
}
