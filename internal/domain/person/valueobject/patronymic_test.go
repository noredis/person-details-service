package person_vo_test

import (
	vo "person-details-service/internal/domain/person/valueobject"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPatronymic(t *testing.T) {
	Convey("Test patronymic", t, func() {
		Convey("Empty patronymic", func() {
			p := ""

			patronymic := vo.NewPatronymic(p)

			So(patronymic, ShouldBeNil)
		})

		Convey("Normal patronymic", func() {
			p := "John"

			patronymic := vo.NewPatronymic(p)

			So(patronymic, ShouldNotBeNil)
			So(patronymic.Value(), ShouldEqual, p)
		})

		Convey("Not equal patronymics", func() {
			p1 := "John"
			p2 := "Doe"

			patronymic1 := vo.NewPatronymic(p1)
			patronymic2 := vo.NewPatronymic(p2)

			So(patronymic1.Equals(*patronymic2), ShouldBeFalse)
		})

		Convey("Equal patronymics", func() {
			p1 := "John"
			p2 := "John"

			patronymic1 := vo.NewPatronymic(p1)
			patronymic2 := vo.NewPatronymic(p2)

			So(patronymic1.Equals(*patronymic2), ShouldBeTrue)
		})
	})
}
