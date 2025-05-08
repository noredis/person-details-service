package person_service_test

import (
	"context"
	"person-details-service/internal/repo/age"
	"person-details-service/internal/repo/gender"
	"person-details-service/internal/repo/nationality"
	"person-details-service/internal/repo/person"
	vo "person-details-service/internal/domain/person/valueobject"
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

				Convey("Find saved person", func() {
					id := person.ID

					person, err := personService.FindPerson(ctx, id)

					So(person, ShouldNotBeNil)
					So(err, ShouldBeNil)
				})

				Convey("Update person", func() {
					updatePersonDTO := dto.UpdatePersonDTO{
						Name:        "John",
						Surname:     "Doe",
						Patronymic:  "John",
						Age:         32,
						Gender:      "male",
						Nationality: "CA",
					}

					id := person.ID

					person, err := personService.UpdatePerson(ctx, id, updatePersonDTO)

					So(person, ShouldNotBeNil)
					So(err, ShouldBeNil)
				})
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

		Convey("Find non-existent person", func() {
			ctx := context.Background()
			id := vo.NewPersonID()

			person, err := personService.FindPerson(ctx, id.Value())

			So(person, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})

		Convey("Find person with invalid id", func() {
			ctx := context.Background()
			id := "asda"

			person, err := personService.FindPerson(ctx, id)

			So(person, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
	})
}
