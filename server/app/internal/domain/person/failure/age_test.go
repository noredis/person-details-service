package person_failure_test

import (
	failure "person-details-service/internal/domain/person/failure"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAgeErrors(t *testing.T) {
	Convey("Test age errors", t, func() {
		Convey("Negative", func() {
			err := failure.ErrAgeIsNegative

			So(err.Error(), ShouldNotBeEmpty)
		})
	})
}
