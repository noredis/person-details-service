//go:build integration
// +build integration

package nationality_infra_test

import (
	"context"
	"log/slog"
	"os"
	"person-details-service/internal/domain/person"
	"person-details-service/internal/domain/person/valueobject"
	infra "person-details-service/internal/infrastructure/nationality"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRealNationalityRepository(t *testing.T) {
	Convey("Test nationality repository", t, func() {
		logger := *slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

		Convey("Valid url", func() {
			repo, err := infra.NewNationalityRepository(logger, "https://api.nationalize.io", 10*time.Second)

			So(repo, ShouldNotBeNil)
			So(err, ShouldBeNil)

			Convey("Get person's nationality", func() {
				ctx := context.Background()

				id := person_vo.NewPersonID()
				name, _ := person_vo.NewName("John")
				surname, _ := person_vo.NewName("Doe")
				patronymic := person_vo.NewPatronymic("")

				johnDoe := person.CreatePerson(id, *name, *surname, patronymic)
				nationality, err := repo.FindOutPersonsNationality(ctx, johnDoe.FullName())

				So(nationality, ShouldNotBeNil)
				So(err, ShouldBeNil)
			})

			Convey("Get nil person's nationality", func() {
				ctx := context.Background()

				id := person_vo.NewPersonID()
				name, _ := person_vo.NewName("asdfsdfsdfasdfasdfasdfsadfsafsdfasdfasdfasdfasdf")
				surname, _ := person_vo.NewName("asdfsdfsdfkljasdlfkasjdflkajsldfkjas;ldkfjasdlf")
				patronymic := person_vo.NewPatronymic("")

				johnDoe := person.CreatePerson(id, *name, *surname, patronymic)
				nationality, err := repo.FindOutPersonsNationality(ctx, johnDoe.FullName())

				So(nationality, ShouldBeNil)
				So(err, ShouldBeNil)
			})
		})

		Convey("Invalid url", func() {
			repo, err := infra.NewNationalityRepository(logger, ":weasdfasdf", 10*time.Second)

			So(repo, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
	})
}
