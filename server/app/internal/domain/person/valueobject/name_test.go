package person_vo_test

import (
	"errors"
	failure "person-details-service/internal/domain/person/failure"
	vo "person-details-service/internal/domain/person/valueobject"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestName(t *testing.T) {
	Convey("Test name", t, func() {
		Convey("Empty name", func() {
			n := ""

			name, err := vo.NewName(n)

			So(name, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(errors.Is(err, failure.ErrNameIsEmpty), ShouldBeTrue)
		})

		Convey("Only spaces", func() {
			n := " "

			name, err := vo.NewName(n)

			So(name, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(errors.Is(err, failure.ErrNameIsEmpty), ShouldBeTrue)
		})

		Convey("Min length name", func() {
			n := "A"

			name, err := vo.NewName(n)

			So(name, ShouldNotBeNil)
			So(name.Value(), ShouldEqual, n)
			So(err, ShouldBeNil)
		})

		Convey("Normal name", func() {
			n := "Doe"

			name, err := vo.NewName(n)

			So(name, ShouldNotBeNil)
			So(name.Value(), ShouldEqual, n)
			So(err, ShouldBeNil)
		})

		Convey("Not equal names", func() {
			n1 := "Doe"
			n2 := "Diamond"

			name1, _ := vo.NewName(n1)
			name2, _ := vo.NewName(n2)

			So(name1.Equals(*name2), ShouldBeFalse)
		})

		Convey("Equal names", func() {
			n1 := "Doe"
			n2 := "Doe"

			name1, _ := vo.NewName(n1)
			name2, _ := vo.NewName(n2)

			So(name1.Equals(*name2), ShouldBeTrue)
		})
	})
}
