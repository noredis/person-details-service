package person_test

import (
	"person-details-service/internal/domain/person"
	vo "person-details-service/internal/domain/person/valueobject"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPerson(t *testing.T) {
	Convey("Test person", t, func() {
		Convey("Create person", func() {
			n := "John"
			s := "Doe"
			p := ""

			name, _ := vo.NewName(n)
			surname, _ := vo.NewName(s)
			patronymic := vo.NewPatronymic(p)

			johnDoe := person.CreatePerson(*name, *surname, patronymic)

			So(johnDoe, ShouldNotBeNil)
			So(johnDoe.Name().Equals(*name), ShouldBeTrue)
			So(johnDoe.Surname().Equals(*surname), ShouldBeTrue)
			So(johnDoe.Patronymic(), ShouldBeNil)
			So(johnDoe.Age(), ShouldBeNil)
			So(johnDoe.Gender(), ShouldBeNil)

			Convey("Get full name", func() {
				So(johnDoe.FullName().Value(), ShouldEqual, "John Doe")
			})

			Convey("Specify age", func() {
				Convey("Non-nil age", func() {
					age, _ := vo.NewAge(38)

					johnDoe.SpecifyAge(age)

					So(johnDoe.Age().Equals(*age), ShouldBeTrue)
				})

				Convey("Nil age", func() {
					age, _ := vo.NewAge(-1)

					johnDoe.SpecifyAge(age)

					So(johnDoe.Age(), ShouldBeNil)
				})
			})

			Convey("Specify gender", func() {
				Convey("Non-nil gender", func() {
					gender, _ := vo.NewGender("male")

					johnDoe.SpecifyGender(gender)

					So(johnDoe.Gender().Equals(*gender), ShouldBeTrue)
				})

				Convey("Nil gender", func() {
					gender, _ := vo.NewGender("asdfa")

					johnDoe.SpecifyGender(gender)

					So(johnDoe.Gender(), ShouldBeNil)
				})
			})
		})
	})
}
