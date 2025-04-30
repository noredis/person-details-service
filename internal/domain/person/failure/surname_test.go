package person_failure_test

import (
	failure "person-details-service/internal/domain/person/failure"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSurnameErrors(t *testing.T) {
	Convey("Test surname errors", t, func() {
		Convey("Empty", func() {
			err := failure.ErrSurnameIsEmpty

			So(err.Error(), ShouldNotBeEmpty)
		})
	})
}
