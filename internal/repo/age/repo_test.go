package age_repo_test

import (
	"context"
	"person-details-service/internal/domain/person"
	"person-details-service/internal/domain/person/valueobject"
	repos "person-details-service/internal/repo/age"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAgeRepo(t *testing.T) {
	Convey("Test age repo", t, func() {
		repo := repos.FakeAgeRepository{}

		Convey("Get age", func() {
			ctx := context.Background()

			id := person_vo.NewPersonID()
			name, _ := person_vo.NewName("John")
			surname, _ := person_vo.NewName("Doe")
			patronymic := person_vo.NewPatronymic("")
			now := time.Now()
			johnDoe := person.CreatePerson(id, *name, *surname, patronymic, now)

			age, err := repo.FindOutPersonsAge(ctx, johnDoe.FullName())

			So(age, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})

		Convey("Get nil age", func() {
			ctx := context.Background()

			id := person_vo.NewPersonID()
			name, _ := person_vo.NewName("John")
			surname, _ := person_vo.NewName("Doe Jr")
			patronymic := person_vo.NewPatronymic("")
			now := time.Now()
			johnDoe := person.CreatePerson(id, *name, *surname, patronymic, now)

			age, err := repo.FindOutPersonsAge(ctx, johnDoe.FullName())

			So(age, ShouldBeNil)
			So(err, ShouldBeNil)
		})
	})
}
