package person_test

import (
	"person-details-service/internal/domain/person"
	vo "person-details-service/internal/domain/person/valueobject"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPerson(t *testing.T) {
	Convey("Test person", t, func() {
		Convey("Create person", func() {
			n := "John"
			s := "Doe"
			p := ""

			id := vo.NewPersonID()
			name, _ := vo.NewName(n)
			surname, _ := vo.NewName(s)
			patronymic := vo.NewPatronymic(p)
			now := time.Now()

			johnDoe := person.CreatePerson(id, *name, *surname, patronymic, now)

			So(johnDoe, ShouldNotBeNil)
			So(johnDoe.ID().Equals(id), ShouldBeTrue)
			So(johnDoe.Name().Equals(*name), ShouldBeTrue)
			So(johnDoe.Surname().Equals(*surname), ShouldBeTrue)
			So(johnDoe.Patronymic(), ShouldBeNil)
			So(johnDoe.Age(), ShouldBeNil)
			So(johnDoe.Gender(), ShouldBeNil)
			So(johnDoe.Nationality(), ShouldBeNil)
			So(johnDoe.CreatedAt().Nanosecond(), ShouldEqual, now.Nanosecond())
			So(johnDoe.UpdatedAt(), ShouldBeNil)

			Convey("Get full name", func() {
				So(johnDoe.FullName().Value(), ShouldEqual, "John Doe")
			})

			Convey("Specify age", func() {
				Convey("Non-nil age", func() {
					age, _ := vo.NewAge(38)

					johnDoe.SpecifyAge(age)

					So(johnDoe.Age().Equals(*age), ShouldBeTrue)
				})

				Convey("Nil age", func() {
					age, _ := vo.NewAge(-1)

					johnDoe.SpecifyAge(age)

					So(johnDoe.Age(), ShouldBeNil)
				})
			})

			Convey("Specify gender", func() {
				Convey("Non-nil gender", func() {
					gender, _ := vo.NewGender("male")

					johnDoe.SpecifyGender(gender)

					So(johnDoe.Gender().Equals(*gender), ShouldBeTrue)
				})

				Convey("Nil gender", func() {
					gender, _ := vo.NewGender("asdfa")

					johnDoe.SpecifyGender(gender)

					So(johnDoe.Gender(), ShouldBeNil)
				})
			})

			Convey("Specify nationality", func() {
				Convey("Non-nil nationality", func() {
					nationality, _ := vo.NewNationality("US")

					johnDoe.SpecifyNationality(nationality)

					So(johnDoe.Nationality().Equals(*nationality), ShouldBeTrue)
				})

				Convey("Nil nationality", func() {
					nationality, _ := vo.NewNationality("dfsadfasdf")

					johnDoe.SpecifyNationality(nationality)

					So(johnDoe.Nationality(), ShouldBeNil)
				})
			})

			Convey("Edit personal information", func() {
				patronymic := vo.NewPatronymic("John")
				age, _ := vo.NewAge(44)
				var gender *vo.Gender
				var nationality *vo.Nationality
				now := time.Now()

				johnDoe.EditPersonalInformation(johnDoe.Name(), johnDoe.Surname(), patronymic, age, gender, nationality, now)

				So(johnDoe.Patronymic().Equals(*patronymic), ShouldBeTrue)
				So(johnDoe.Age().Equals(*age), ShouldBeTrue)
				So(johnDoe.Gender(), ShouldBeNil)
				So(johnDoe.Nationality(), ShouldBeNil)
				So(johnDoe.UpdatedAt().UnixNano(), ShouldEqual, now.UnixNano())
			})
		})

		Convey("Restore person", func() {
			id := vo.NewPersonID()
			name, _ := vo.NewName("John")
			surname, _ := vo.NewName("Doe")
			patronymic := vo.NewPatronymic("")
			age, _ := vo.NewAge(38)
			gender, _ := vo.NewGender("male")
			nationality, _ := vo.NewNationality("CA")
			createdAt := time.Now()
			var updatedAt *time.Time

			johnDoe := person.RestorePerson(id, *name, *surname, patronymic, age, gender, nationality, createdAt, updatedAt)

			So(johnDoe.ID().Equals(id), ShouldBeTrue)
			So(johnDoe.Name().Equals(*name), ShouldBeTrue)
			So(johnDoe.Surname().Equals(*surname), ShouldBeTrue)
			So(johnDoe.Patronymic(), ShouldBeNil)
			So(johnDoe.Age().Equals(*age), ShouldBeTrue)
			So(johnDoe.Gender().Equals(*gender), ShouldBeTrue)
			So(johnDoe.Nationality().Equals(*nationality), ShouldBeTrue)
			So(johnDoe.CreatedAt().Nanosecond(), ShouldEqual, createdAt.Nanosecond())
			So(johnDoe.UpdatedAt(), ShouldBeNil)
		})
	})
}
