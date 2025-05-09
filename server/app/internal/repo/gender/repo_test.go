package gender_repo_test

import (
	"context"
	"person-details-service/internal/domain/person"
	"person-details-service/internal/domain/person/valueobject"
	repos "person-details-service/internal/repo/gender"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGenderRepo(t *testing.T) {
	Convey("Test gender repo", t, func() {
		repo := repos.FakeGenderRepository{}

		Convey("Get gender", func() {
			ctx := context.Background()

			id := person_vo.NewPersonID()
			name, _ := person_vo.NewName("John")
			surname, _ := person_vo.NewName("Doe")
			patronymic := person_vo.NewPatronymic("")
			now := time.Now()
			johnDoe := person.CreatePerson(id, *name, *surname, patronymic, now)

			gender, err := repo.FindOutPersonsGender(ctx, johnDoe.FullName())

			So(gender, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})

		Convey("Get nil gender", func() {
			ctx := context.Background()

			id := person_vo.NewPersonID()
			name, _ := person_vo.NewName("John")
			surname, _ := person_vo.NewName("Doe")
			patronymic := person_vo.NewPatronymic("John")
			now := time.Now()
			johnDoe := person.CreatePerson(id, *name, *surname, patronymic, now)

			gender, err := repo.FindOutPersonsGender(ctx, johnDoe.FullName())

			So(gender, ShouldBeNil)
			So(err, ShouldBeNil)
		})

		Convey("Get error gender", func() {
			ctx := context.Background()

			id := person_vo.NewPersonID()
			name, _ := person_vo.NewName("John")
			surname, _ := person_vo.NewName("Doe")
			patronymic := person_vo.NewPatronymic("2")
			now := time.Now()
			johnDoe := person.CreatePerson(id, *name, *surname, patronymic, now)

			gender, err := repo.FindOutPersonsGender(ctx, johnDoe.FullName())

			So(gender, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
	})
}
