package person_vo_test

import (
	"errors"
	failure "person-details-service/internal/domain/person/failure"
	vo "person-details-service/internal/domain/person/valueobject"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGender(t *testing.T) {
	Convey("Test gender", t, func() {
		Convey("Male", func() {
			g := "male"

			gender, err := vo.NewGender(g)

			So(gender, ShouldNotBeNil)
			So(gender.Value(), ShouldEqual, g)
			So(gender.IsMale(), ShouldBeTrue)
			So(gender.IsFemale(), ShouldBeFalse)
			So(err, ShouldBeNil)
		})

		Convey("Female", func() {
			g := "female"

			gender, err := vo.NewGender(g)

			So(gender, ShouldNotBeNil)
			So(gender.Value(), ShouldEqual, g)
			So(gender.IsMale(), ShouldBeFalse)
			So(gender.IsFemale(), ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		Convey("Empty", func() {
			g := ""

			gender, err := vo.NewGender(g)

			So(gender, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(errors.Is(err, failure.ErrUnexpectedGender), ShouldBeTrue)
		})

		Convey("Unexpected", func() {
			g := "asdfasdf"

			gender, err := vo.NewGender(g)

			So(gender, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(errors.Is(err, failure.ErrUnexpectedGender), ShouldBeTrue)
		})

		Convey("Not equal genders", func() {
			g1 := "male"
			g2 := "female"

			gender1, _ := vo.NewGender(g1)
			gender2, _ := vo.NewGender(g2)

			So(gender1.Equals(*gender2), ShouldBeFalse)
		})

		Convey("Equal genders", func() {
			g1 := "male"
			g2 := "male"

			gender1, _ := vo.NewGender(g1)
			gender2, _ := vo.NewGender(g2)

			So(gender1.Equals(*gender2), ShouldBeTrue)
		})
	})
}
