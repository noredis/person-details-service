package person_vo_test

import (
	"errors"
	failure "person-details-service/internal/domain/person/failure"
	vo "person-details-service/internal/domain/person/valueobject"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNationality(t *testing.T) {
	Convey("Test nationality", t, func() {
		Convey("Empty nationality", func() {
			n := ""

			nationality, err := vo.NewNationality(n)

			So(nationality, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(errors.Is(err, failure.ErrUnexpectedNationality), ShouldBeTrue)
		})

		Convey("Normal nationality", func() {
			n := "RU"

			nationality, err := vo.NewNationality(n)

			So(nationality, ShouldNotBeNil)
			So(nationality.Value(), ShouldEqual, n)
			So(err, ShouldBeNil)
		})

		Convey("Unexpected nationality", func() {
			n := "d"

			nationality, err := vo.NewNationality(n)

			So(nationality, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(errors.Is(err, failure.ErrUnexpectedNationality), ShouldBeTrue)
		})

		Convey("Not equal nationalities", func() {
			n1 := "RU"
			n2 := "KZ"

			nationality1, _ := vo.NewNationality(n1)
			nationality2, _ := vo.NewNationality(n2)

			So(nationality1.Equals(*nationality2), ShouldBeFalse)
		})

		Convey("Equal nationalities", func() {
			n1 := "RU"
			n2 := "RU"

			nationality1, _ := vo.NewNationality(n1)
			nationality2, _ := vo.NewNationality(n2)

			So(nationality1.Equals(*nationality2), ShouldBeTrue)
		})
	})
}
