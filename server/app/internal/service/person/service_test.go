package person_service_test

import (
	"context"
	"person-details-service/internal/repo/age"
	"person-details-service/internal/repo/gender"
	"person-details-service/internal/repo/nationality"
	"person-details-service/internal/repo/person"
	service "person-details-service/internal/service/person"
	dto "person-details-service/internal/service/person/dto"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPersonService(t *testing.T) {
	Convey("Test person service", t, func() {
		ageRepo := age_repo.FakeAgeRepository{}
		genderRepo := gender_repo.FakeGenderRepository{}
		nationalityRepo := nationality_repo.FakeNationalityRepository{}
		personRepo := person_repo.NewFakePersonRepository()
		personService := service.NewPersonService(ageRepo, genderRepo, nationalityRepo, personRepo)

		Convey("Create person", func() {
			Convey("Filled in completely person", func() {
				ctx := context.Background()
				createPersonDTO := dto.CreatePersonDTO{
					Name:       "John",
					Surname:    "Doe",
					Patronymic: "John",
				}

				person, err := personService.CreatePerson(ctx, createPersonDTO)

				So(person, ShouldNotBeNil)
				So(person.Name, ShouldEqual, createPersonDTO.Name)
				So(person.Surname, ShouldEqual, createPersonDTO.Surname)
				So(*person.Patronymic, ShouldEqual, createPersonDTO.Patronymic)
				So(err, ShouldBeNil)
			})

			Convey("Filled in minimally person", func() {
				ctx := context.Background()
				createPersonDTO := dto.CreatePersonDTO{
					Name:    "John",
					Surname: "Doe",
				}

				person, err := personService.CreatePerson(ctx, createPersonDTO)

				So(person, ShouldNotBeNil)
				So(person.Name, ShouldEqual, createPersonDTO.Name)
				So(person.Surname, ShouldEqual, createPersonDTO.Surname)
				So(err, ShouldBeNil)
			})

			Convey("Empty name", func() {
				ctx := context.Background()
				createPersonDTO := dto.CreatePersonDTO{
					Name:    "",
					Surname: "Doe",
				}

				person, err := personService.CreatePerson(ctx, createPersonDTO)

				So(person, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})

			Convey("Empty surname", func() {
				ctx := context.Background()
				createPersonDTO := dto.CreatePersonDTO{
					Name:    "John",
					Surname: "",
				}

				person, err := personService.CreatePerson(ctx, createPersonDTO)

				So(person, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})

			Convey("Nil name", func() {
				ctx := context.Background()
				createPersonDTO := dto.CreatePersonDTO{
					Surname: "Doe",
				}

				person, err := personService.CreatePerson(ctx, createPersonDTO)

				So(person, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})

			Convey("Nil surname", func() {
				ctx := context.Background()
				createPersonDTO := dto.CreatePersonDTO{
					Name: "John",
				}

				person, err := personService.CreatePerson(ctx, createPersonDTO)

				So(person, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
	})
}
