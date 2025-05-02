package person_failure_test

import (
	failure "person-details-service/internal/domain/person/failure"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPersonIDErrors(t *testing.T) {
	Convey("Test person id errors", t, func() {
		Convey("Parse", func() {
			err := failure.ErrParsePersonID

			So(err.Error(), ShouldNotBeEmpty)
		})
	})
}
