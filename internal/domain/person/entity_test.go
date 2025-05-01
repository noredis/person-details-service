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

			Convey("Get full name", func() {
				So(johnDoe.FullName().Value(), ShouldEqual, "John Doe")
			})
		})
	})
}
