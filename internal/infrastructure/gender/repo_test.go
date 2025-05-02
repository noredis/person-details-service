//go:build integration
// +build integration

package gender_infra_test

import (
	"context"
	"log/slog"
	"os"
	"person-details-service/internal/domain/person"
	"person-details-service/internal/domain/person/valueobject"
	infra "person-details-service/internal/infrastructure/gender"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRealGenderRepository(t *testing.T) {
	Convey("Test gender repository", t, func() {
		logger := *slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

		Convey("Valid url", func() {
			repo, err := infra.NewGenderRepository(logger, "https://api.genderize.io", 10*time.Second)

			So(repo, ShouldNotBeNil)
			So(err, ShouldBeNil)

			Convey("Get person's gender", func() {
				ctx := context.Background()

				id := person_vo.NewPersonID()
				name, _ := person_vo.NewName("John")
				surname, _ := person_vo.NewName("Doe")
				patronymic := person_vo.NewPatronymic("")

				johnDoe := person.CreatePerson(id, *name, *surname, patronymic)
				gender, err := repo.FindOutPersonsGender(ctx, johnDoe.FullName())

				So(gender, ShouldNotBeNil)
				So(err, ShouldBeNil)
			})

			Convey("Get nil person's gender", func() {
				ctx := context.Background()

				id := person_vo.NewPersonID()
				name, _ := person_vo.NewName("asdfsdfsdfasdfasdfasdfsadfsafsdfasdfasdfasdfasdf")
				surname, _ := person_vo.NewName("asdfsdfsdfkljasdlfkasjdflkajsldfkjas;ldkfjasdlf")
				patronymic := person_vo.NewPatronymic("")

				johnDoe := person.CreatePerson(id, *name, *surname, patronymic)
				gender, err := repo.FindOutPersonsGender(ctx, johnDoe.FullName())

				So(gender, ShouldBeNil)
				So(err, ShouldBeNil)
			})
		})

		Convey("Invalid url", func() {
			repo, err := infra.NewGenderRepository(logger, ":weasdfasdf", 10*time.Second)

			So(repo, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
	})
}
