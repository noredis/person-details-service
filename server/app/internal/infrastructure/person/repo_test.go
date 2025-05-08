//go:build integration
// +build integration

package person_infra_test

import (
	"context"
	"person-details-service/internal/domain/person"
	vo "person-details-service/internal/domain/person/valueobject"
	repos "person-details-service/internal/infrastructure/person"
	"person-details-service/pkg/testingpg"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRealPersonRepository(t *testing.T) {
	Convey("Test real person repository", t, func() {
		Convey("Save person", func() {
			ctx := context.Background()
			db := testingpg.NewWithIsolatedDatabase(t)
			repo := repos.NewPersonRepository(db.DB())

			id := vo.NewPersonID()
			name, _ := vo.NewName("John")
			surname, _ := vo.NewName("Doe")
			patronymic := vo.NewPatronymic("")
			age, _ := vo.NewAge(38)
			gender, _ := vo.NewGender("male")
			nationality, _ := vo.NewNationality("US")
			now := time.Now()

			johnDoe := person.CreatePerson(id, *name, *surname, patronymic, now)
			johnDoe.SpecifyAge(age)
			johnDoe.SpecifyGender(gender)
			johnDoe.SpecifyNationality(nationality)

			err := repo.SavePerson(ctx, *johnDoe)

			So(err, ShouldBeNil)

			Convey("Get saved person", func() {
				restoredJohnDoe, err := repo.GetPersonByID(ctx, id)

				So(restoredJohnDoe, ShouldNotBeNil)
				So(restoredJohnDoe.ID().Equals(id), ShouldBeTrue)
				So(restoredJohnDoe.Name().Equals(*name), ShouldBeTrue)
				So(restoredJohnDoe.Surname().Equals(*surname), ShouldBeTrue)
				So(restoredJohnDoe.Patronymic(), ShouldBeNil)
				So(restoredJohnDoe.Age().Equals(*age), ShouldBeTrue)
				So(restoredJohnDoe.Gender().Equals(*gender), ShouldBeTrue)
				So(restoredJohnDoe.Nationality().Equals(*nationality), ShouldBeTrue)
				So(restoredJohnDoe.UpdatedAt(), ShouldEqual, johnDoe.UpdatedAt())
				So(err, ShouldBeNil)
			})

			Convey("Update person", func() {
				patronymic := vo.NewPatronymic("John")
				now = time.Now()

				johnDoe.EditPersonalInformation(
					johnDoe.Name(),
					johnDoe.Surname(),
					patronymic,
					johnDoe.Age(),
					johnDoe.Gender(),
					johnDoe.Nationality(),
					now,
				)

				err = repo.UpdatePerson(ctx, *johnDoe)

				So(err, ShouldBeNil)

				restoredJohnDoe, _ := repo.GetPersonByID(ctx, id)
				So(restoredJohnDoe.Patronymic(), ShouldNotBeNil)
				So(restoredJohnDoe.Patronymic().Equals(*patronymic), ShouldBeTrue)

				Convey("Delete person", func() {
					err = repo.DeletePerson(ctx, johnDoe.ID())

					So(err, ShouldBeNil)
				})
			})
		})

		Convey("Get non-existent person", func() {
			ctx := context.Background()
			db := testingpg.NewWithIsolatedDatabase(t)
			repo := repos.NewPersonRepository(db.DB())

			id := vo.NewPersonID()

			who, err := repo.GetPersonByID(ctx, id)

			So(who, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})

		Convey("Delete non-existent person", func() {
			ctx := context.Background()
			db := testingpg.NewWithIsolatedDatabase(t)
			repo := repos.NewPersonRepository(db.DB())

			id := vo.NewPersonID()

			err := repo.DeletePerson(ctx, id)

			So(err, ShouldBeNil)
		})
	})
}
