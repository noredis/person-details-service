package age_repo_test

import (
	"context"
	"person-details-service/internal/domain/person"
	vo "person-details-service/internal/domain/person/valueobject"
	repo "person-details-service/internal/repo/age"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAgeRepo(t *testing.T) {
	Convey("Test age repo", t, func() {
		repo := repo.FakeAgeRepository{}

		Convey("Get age", func() {
			ctx := context.Background()

			name, _ := vo.NewName("John")
			surname, _ := vo.NewName("Doe")
			patronymic := vo.NewPatronymic("")
			johnDoe := person.CreatePerson(*name, *surname, patronymic)

			age, err := repo.FindOutPersonsAge(ctx, *johnDoe)

			So(age, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})

		Convey("Get nil age", func() {
			ctx := context.Background()

			name, _ := vo.NewName("John")
			surname, _ := vo.NewName("Doe Jr")
			patronymic := vo.NewPatronymic("")
			johnDoe := person.CreatePerson(*name, *surname, patronymic)

			age, err := repo.FindOutPersonsAge(ctx, *johnDoe)

			So(age, ShouldBeNil)
			So(err, ShouldBeNil)
		})
	})
}
