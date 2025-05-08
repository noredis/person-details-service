package nationality_repo_test

import (
	"context"
	"person-details-service/internal/domain/person"
	"person-details-service/internal/domain/person/valueobject"
	repos "person-details-service/internal/repo/nationality"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNationalityRepo(t *testing.T) {
	Convey("Test nationality repo", t, func() {
		repo := repos.FakeNationalityRepository{}

		Convey("Get nationality", func() {
			ctx := context.Background()

			id := person_vo.NewPersonID()
			name, _ := person_vo.NewName("John")
			surname, _ := person_vo.NewName("Doe")
			patronymic := person_vo.NewPatronymic("")
			now := time.Now()
			johnDoe := person.CreatePerson(id, *name, *surname, patronymic, now)

			nationality, err := repo.FindOutPersonsNationality(ctx, johnDoe.FullName())

			So(nationality, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})

		Convey("Get nil nationality", func() {
			ctx := context.Background()

			id := person_vo.NewPersonID()
			name, _ := person_vo.NewName("John")
			surname, _ := person_vo.NewName("Doe")
			patronymic := person_vo.NewPatronymic("John")
			now := time.Now()
			johnDoe := person.CreatePerson(id, *name, *surname, patronymic, now)

			nationality, err := repo.FindOutPersonsNationality(ctx, johnDoe.FullName())

			So(nationality, ShouldBeNil)
			So(err, ShouldBeNil)
		})

		Convey("Get error nationality", func() {
			ctx := context.Background()

			id := person_vo.NewPersonID()
			name, _ := person_vo.NewName("John")
			surname, _ := person_vo.NewName("Doe")
			patronymic := person_vo.NewPatronymic("3")
			now := time.Now()
			johnDoe := person.CreatePerson(id, *name, *surname, patronymic, now)

			nationality, err := repo.FindOutPersonsNationality(ctx, johnDoe.FullName())

			So(nationality, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
	})
}
