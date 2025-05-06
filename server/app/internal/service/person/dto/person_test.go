package person_dto_test

import (
	"person-details-service/internal/domain/person"
	vo "person-details-service/internal/domain/person/valueobject"
	dto "person-details-service/internal/service/person/dto"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPersonDTO(t *testing.T) {
	Convey("Test person DTO", t, func() {
		Convey("Map from person", func() {
			Convey("Map filled in completely", func() {
				id := vo.NewPersonID()
				name, _ := vo.NewName("John")
				surname, _ := vo.NewName("Doe")
				patronymic := vo.NewPatronymic("John")
				age, _ := vo.NewAge(28)
				gender, _ := vo.NewGender("male")
				nationality, _ := vo.NewNationality("US")
				now := time.Now()

				johnDoe := person.CreatePerson(id, *name, *surname, patronymic, now)

				johnDoe.SpecifyAge(age)
				johnDoe.SpecifyGender(gender)
				johnDoe.SpecifyNationality(nationality)

				pDTO := dto.MapFromPerson(*johnDoe)

				So(pDTO.ID, ShouldEqual, johnDoe.ID().Value())
				So(pDTO.Name, ShouldEqual, johnDoe.Name().Value())
				So(pDTO.Surname, ShouldEqual, johnDoe.Surname().Value())
				So(*pDTO.Patronymic, ShouldEqual, johnDoe.Patronymic().Value())
				So(*pDTO.Age, ShouldEqual, johnDoe.Age().Value())
				So(*pDTO.Gender, ShouldEqual, johnDoe.Gender().Value())
				So(*pDTO.Nationality, ShouldEqual, johnDoe.Nationality().Value())
			})
		})
	})
}
