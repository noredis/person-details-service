package person_vo_test

import (
	vo "person-details-service/internal/domain/person/valueobject"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFullName(t *testing.T) {
	Convey("Test full name", t, func() {
		Convey("Full name without patronymic", func() {
			name, _ := vo.NewName("John")
			surname, _ := vo.NewName("Doe")
			patronymic := vo.NewPatronymic("")

			fullName := vo.NewFullName(*name, *surname, patronymic)

			So(fullName.Value(), ShouldEqual, "John Doe")
		})

		Convey("Full name with patronymic", func() {
			name, _ := vo.NewName("John")
			surname, _ := vo.NewName("Doe")
			patronymic := vo.NewPatronymic("John")

			fullName := vo.NewFullName(*name, *surname, patronymic)

			So(fullName.Value(), ShouldEqual, "John Doe John")
		})

		Convey("Not equal full names", func() {
			name1, _ := vo.NewName("John")
			surname1, _ := vo.NewName("Doe")
			patronymic1 := vo.NewPatronymic("")

			name2, _ := vo.NewName("John")
			surname2, _ := vo.NewName("Doe")
			patronymic2 := vo.NewPatronymic("John")

			fullName1 := vo.NewFullName(*name1, *surname1, patronymic1)
			fullName2 := vo.NewFullName(*name2, *surname2, patronymic2)

			So(fullName1.Equals(fullName2), ShouldBeFalse)
		})

		Convey("Equal full names", func() {
			name1, _ := vo.NewName("John")
			surname1, _ := vo.NewName("Doe")
			patronymic1 := vo.NewPatronymic("")

			name2, _ := vo.NewName("John")
			surname2, _ := vo.NewName("Doe")
			patronymic2 := vo.NewPatronymic("")

			fullName1 := vo.NewFullName(*name1, *surname1, patronymic1)
			fullName2 := vo.NewFullName(*name2, *surname2, patronymic2)

			So(fullName1.Equals(fullName2), ShouldBeTrue)
		})
	})
}
