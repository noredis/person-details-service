package person_failure_test

import (
	failure "person-details-service/internal/domain/person/failure"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGenderErrors(t *testing.T) {
	Convey("Test gender errors", t, func() {
		Convey("Unexpected", func() {
			err := failure.ErrUnexpectedGender

			So(err.Error(), ShouldNotBeEmpty)
		})
	})
}
