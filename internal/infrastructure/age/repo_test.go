//go:build integration
// +build integration

package age_infra_test

import (
	"context"
	"log/slog"
	"os"
	"person-details-service/internal/domain/person"
	"person-details-service/internal/domain/person/valueobject"
	infra "person-details-service/internal/infrastructure/age"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRealAgeRepository(t *testing.T) {
	Convey("Test age repository", t, func() {
		logger := *slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

		Convey("Valid url", func() {
			repo, err := infra.NewAgeRepository(logger, "https://api.agify.io", 10*time.Second)

			So(repo, ShouldNotBeNil)
			So(err, ShouldBeNil)

			Convey("Get person's age", func() {
				ctx := context.Background()

				name, _ := person_vo.NewName("John")
				surname, _ := person_vo.NewName("Doe")
				patronymic := person_vo.NewPatronymic("")

				johnDoe := person.CreatePerson(*name, *surname, patronymic)
				age, err := repo.FindOutPersonsAge(ctx, *johnDoe)

				So(age, ShouldNotBeNil)
				So(err, ShouldBeNil)
			})

			Convey("Get nil person's age", func() {
				ctx := context.Background()

				name, _ := person_vo.NewName("asdfsdfsdfasdfasdfasdfsadfsafsdfasdfasdfasdfasdf")
				surname, _ := person_vo.NewName("asdfsdfsdfkljasdlfkasjdflkajsldfkjas;ldkfjasdlf")
				patronymic := person_vo.NewPatronymic("")

				johnDoe := person.CreatePerson(*name, *surname, patronymic)
				age, err := repo.FindOutPersonsAge(ctx, *johnDoe)

				So(age, ShouldBeNil)
				So(err, ShouldBeNil)
			})
		})

		Convey("Invalid url", func() {
			repo, err := infra.NewAgeRepository(logger, ":weasdfasdf", 10*time.Second)

			So(repo, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
	})
}
