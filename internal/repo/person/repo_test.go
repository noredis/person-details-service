package person_repo_test

import (
	"context"
	"person-details-service/internal/domain/person"
	vo "person-details-service/internal/domain/person/valueobject"
	repos "person-details-service/internal/repo/person"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPersonRepo(t *testing.T) {
	Convey("Test person repo", t, func() {
		repo := repos.NewFakePersonRepository()

		Convey("Save person", func() {
			ctx := context.Background()

			name, _ := vo.NewName("John")
			surname, _ := vo.NewName("Doe")
			patronymic := vo.NewPatronymic("")
			age, _ := vo.NewAge(28)
			gender, _ := vo.NewGender("male")
			nationality, _ := vo.NewNationality("asdfas")

			johnDoe := person.CreatePerson(*name, *surname, patronymic)

			johnDoe.SpecifyAge(age)
			johnDoe.SpecifyGender(gender)
			johnDoe.SpecifyNationality(nationality)

			err := repo.SavePerson(ctx, *johnDoe)

			So(err, ShouldBeNil)
		})
	})
}
