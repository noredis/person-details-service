package person_vo_test

import (
	"errors"
	failure "person-details-service/internal/domain/person/failure"
	vo "person-details-service/internal/domain/person/valueobject"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAge(t *testing.T) {
	Convey("Test age", t, func() {
		Convey("Negative age", func() {
			a := -1

			age, err := vo.NewAge(a)

			So(age, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(errors.Is(err, failure.ErrAgeIsNegative), ShouldBeTrue)
		})

		Convey("Zero age", func() {
			a := 0

			age, err := vo.NewAge(a)

			So(age, ShouldNotBeNil)
			So(age.Value(), ShouldEqual, a)
			So(err, ShouldBeNil)
		})

		Convey("Normal age", func() {
			a := 38

			age, err := vo.NewAge(a)

			So(age, ShouldNotBeNil)
			So(age.Value(), ShouldEqual, a)
			So(err, ShouldBeNil)
		})

		Convey("Not equal ages", func() {
			a1 := 39
			a2 := 21

			age1, _ := vo.NewAge(a1)
			age2, _ := vo.NewAge(a2)

			So(age1.Equals(*age2), ShouldBeFalse)
		})

		Convey("Equal ages", func() {
			a1 := 20
			a2 := 20

			age1, _ := vo.NewAge(a1)
			age2, _ := vo.NewAge(a2)

			So(age1.Equals(*age2), ShouldBeTrue)
		})
	})
}
