package person_vo_test

import (
	"errors"
	failure "person-details-service/internal/domain/person/failure"
	vo "person-details-service/internal/domain/person/valueobject"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPersonID(t *testing.T) {
	Convey("Test person id", t, func() {
		Convey("New person id", func() {
			id := vo.NewPersonID()

			So(id.Value(), ShouldNotBeEmpty)
		})

		Convey("Parse invalid person id", func() {
			i := ""

			id, err := vo.ParsePersonID(i)

			So(id, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(errors.Is(err, failure.ErrParsePersonID), ShouldBeTrue)
		})

		Convey("Parse valid person id", func() {
			id := vo.NewPersonID()

			parsedID, err := vo.ParsePersonID(id.Value())

			So(parsedID, ShouldNotBeNil)
			So(parsedID.Value(), ShouldEqual, id.Value())
			So(err, ShouldBeNil)
		})

		Convey("Not equal person ids", func() {
			id1 := vo.NewPersonID()
			id2 := vo.NewPersonID()

			So(id1.Equals(id2), ShouldBeFalse)
		})

		Convey("Equal person ids", func() {
			id1 := vo.NewPersonID()
			id2, _ := vo.ParsePersonID(id1.Value())

			So(id1.Equals(*id2), ShouldBeTrue)
		})
	})
}
