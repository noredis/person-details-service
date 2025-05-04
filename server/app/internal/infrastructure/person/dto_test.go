package person_infra_test

import (
	"person-details-service/internal/domain/person"
	vo "person-details-service/internal/domain/person/valueobject"
	infra "person-details-service/internal/infrastructure/person"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPersonDTO(t *testing.T) {
	Convey("Test person dto", t, func() {
		Convey("Map person to dto", func() {
			Convey("Map filled in completely", func() {
				id := vo.NewPersonID()
				name, _ := vo.NewName("John")
				surname, _ := vo.NewName("Doe")
				patronymic := vo.NewPatronymic("John")
				age, _ := vo.NewAge(38)
				gender, _ := vo.NewGender("male")
				nationality, _ := vo.NewNationality("US")
				now := time.Now()

				johnDoe := person.CreatePerson(id, *name, *surname, patronymic, now)
				johnDoe.SpecifyAge(age)
				johnDoe.SpecifyGender(gender)
				johnDoe.SpecifyNationality(nationality)

				dto := infra.MapPersonToDTO(*johnDoe)

				So(dto.ID, ShouldEqual, johnDoe.ID().Value())
				So(dto.Name, ShouldEqual, johnDoe.Name().Value())
				So(dto.Surname, ShouldEqual, johnDoe.Surname().Value())
				So(dto.Patronymic.String, ShouldEqual, johnDoe.Patronymic().Value())
				So(dto.Patronymic.Valid, ShouldBeTrue)
				So(dto.Age.Int16, ShouldEqual, johnDoe.Age().Value())
				So(dto.Age.Valid, ShouldBeTrue)
				So(dto.Gender.String, ShouldEqual, johnDoe.Gender().Value())
				So(dto.Gender.Valid, ShouldBeTrue)
				So(dto.Nationality.String, ShouldEqual, johnDoe.Nationality().Value())
				So(dto.Nationality.Valid, ShouldBeTrue)
				So(dto.CreatedAt, ShouldEqual, johnDoe.CreatedAt())
				So(dto.UpdatedAt, ShouldEqual, johnDoe.UpdatedAt())
			})

			Convey("Map filled in minimally", func() {
				id := vo.NewPersonID()
				name, _ := vo.NewName("John")
				surname, _ := vo.NewName("Doe")
				patronymic := vo.NewPatronymic("")
				now := time.Now()

				johnDoe := person.CreatePerson(id, *name, *surname, patronymic, now)

				dto := infra.MapPersonToDTO(*johnDoe)

				So(dto.ID, ShouldEqual, johnDoe.ID().Value())
				So(dto.Name, ShouldEqual, johnDoe.Name().Value())
				So(dto.Surname, ShouldEqual, johnDoe.Surname().Value())
				So(dto.Patronymic.Valid, ShouldBeFalse)
				So(dto.Age.Valid, ShouldBeFalse)
				So(dto.Gender.Valid, ShouldBeFalse)
				So(dto.Nationality.Valid, ShouldBeFalse)
				So(dto.CreatedAt, ShouldEqual, johnDoe.CreatedAt())
				So(dto.UpdatedAt, ShouldEqual, johnDoe.UpdatedAt())
			})
		})
	})
}
