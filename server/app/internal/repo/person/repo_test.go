package person_repo_test

import (
	"context"
	"person-details-service/internal/domain/person"
	vo "person-details-service/internal/domain/person/valueobject"
	repos "person-details-service/internal/repo/person"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPersonRepo(t *testing.T) {
	Convey("Test person repo", t, func() {
		repo := repos.NewFakePersonRepository()

		Convey("Save person", func() {
			ctx := context.Background()

			id := vo.NewPersonID()
			name, _ := vo.NewName("John")
			surname, _ := vo.NewName("Doe")
			patronymic := vo.NewPatronymic("")
			age, _ := vo.NewAge(28)
			gender, _ := vo.NewGender("male")
			nationality, _ := vo.NewNationality("US")
			now := time.Now()

			johnDoe := person.CreatePerson(id, *name, *surname, patronymic, now)

			johnDoe.SpecifyAge(age)
			johnDoe.SpecifyGender(gender)
			johnDoe.SpecifyNationality(nationality)

			err := repo.SavePerson(ctx, *johnDoe)

			So(err, ShouldBeNil)

			Convey("Get saved person matches all filters", func() {
				filterOptions := repos.FilterOptions{Age: age, Gender: gender}

				persons, err := repo.GetPersons(ctx, filterOptions)

				So(len(persons), ShouldEqual, 1)
				So(err, ShouldBeNil)
			})

			Convey("Get saved person not matches age filters", func() {
				age, _ = vo.NewAge(30)
				filterOptions := repos.FilterOptions{Age: age, Gender: gender}

				persons, err := repo.GetPersons(ctx, filterOptions)

				So(len(persons), ShouldEqual, 0)
				So(err, ShouldBeNil)
			})

			Convey("Get saved person not matches gender filters", func() {
				gender, _ = vo.NewGender("female")
				filterOptions := repos.FilterOptions{Age: age, Gender: gender}

				persons, err := repo.GetPersons(ctx, filterOptions)

				So(len(persons), ShouldEqual, 0)
				So(err, ShouldBeNil)
			})

			Convey("Get saved person not matches nationality filters", func() {
				nationality, _ = vo.NewNationality("CA")
				filterOptions := repos.FilterOptions{Age: age, Gender: gender, Nationality: nationality}

				persons, err := repo.GetPersons(ctx, filterOptions)

				So(len(persons), ShouldEqual, 0)
				So(err, ShouldBeNil)
			})

			Convey("Get saved person", func() {
				restoredJohnDoe, err := repo.GetPersonByID(ctx, id)

				So(restoredJohnDoe.ID().Equals(id), ShouldBeTrue)
				So(err, ShouldBeNil)
			})

			Convey("Update person", func() {
				err = repo.UpdatePerson(ctx, *johnDoe)

				So(err, ShouldBeNil)
			})

			Convey("Delete person", func() {
				err = repo.DeletePerson(ctx, johnDoe.ID())

				So(err, ShouldBeNil)
			})
		})

		Convey("Update non-existent person", func() {
			ctx := context.Background()

			id := vo.NewPersonID()
			name, _ := vo.NewName("John")
			surname, _ := vo.NewName("Doe")
			patronymic := vo.NewPatronymic("")
			now := time.Now()

			johnDoe := person.CreatePerson(id, *name, *surname, patronymic, now)

			err := repo.UpdatePerson(ctx, *johnDoe)

			So(err, ShouldBeNil)
		})

		Convey("Get non-existent person", func() {
			ctx := context.Background()

			id := vo.NewPersonID()

			who, err := repo.GetPersonByID(ctx, id)

			So(who, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})

		Convey("Delete non-existent person", func() {
			ctx := context.Background()

			id := vo.NewPersonID()

			err := repo.DeletePerson(ctx, id)

			So(err, ShouldBeNil)
		})
	})
}
