package person_vo_test

import (
	"errors"
	failure "person-details-service/internal/domain/person/failure"
	vo "person-details-service/internal/domain/person/valueobject"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSurname(t *testing.T) {
	Convey("Test surname", t, func() {
		Convey("Empty surname", func() {
			s := ""

			surname, err := vo.NewSurname(s)

			So(surname, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(errors.Is(err, failure.ErrSurnameIsEmpty), ShouldBeTrue)
		})

		Convey("Min length surname", func() {
			s := "A"

			surname, err := vo.NewSurname(s)

			So(surname, ShouldNotBeNil)
			So(surname.Value(), ShouldEqual, s)
			So(err, ShouldBeNil)
		})

		Convey("Normal surname", func() {
			s := "Doe"

			surname, err := vo.NewSurname(s)

			So(surname, ShouldNotBeNil)
			So(surname.Value(), ShouldEqual, s)
			So(err, ShouldBeNil)
		})

		Convey("Not equal surnames", func() {
			s1 := "Doe"
			s2 := "Diamond"

			surname1, _ := vo.NewSurname(s1)
			surname2, _ := vo.NewSurname(s2)

			So(surname1.Equals(*surname2), ShouldBeFalse)
		})

		Convey("Equal surnames", func() {
			s1 := "Doe"
			s2 := "Doe"

			surname1, _ := vo.NewSurname(s1)
			surname2, _ := vo.NewSurname(s2)

			So(surname1.Equals(*surname2), ShouldBeTrue)
		})
	})
}
