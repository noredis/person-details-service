package person_failure_test

import (
	failure "person-details-service/internal/domain/person/failure"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNationalityErrors(t *testing.T) {
	Convey("Test nationality errors", t, func() {
		Convey("Unexpected", func() {
			err := failure.ErrUnexpectedNationality

			So(err.Error(), ShouldNotBeEmpty)
		})
	})
}
