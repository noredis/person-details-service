package person_app_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	app "person-details-service/internal/app/person"
	"person-details-service/internal/repo/age"
	"person-details-service/internal/repo/gender"
	"person-details-service/internal/repo/nationality"
	"person-details-service/internal/repo/person"
	service "person-details-service/internal/service/person"
	dto "person-details-service/internal/service/person/dto"
	"testing"

	"github.com/julienschmidt/httprouter"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPersonHandler(t *testing.T) {
	Convey("Test person handler", t, func() {
		ctx := context.Background()
		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

		ageRepo := age_repo.FakeAgeRepository{}
		genderRepo := gender_repo.FakeGenderRepository{}
		nationalityRepo := nationality_repo.FakeNationalityRepository{}
		personRepo := person_repo.NewFakePersonRepository()

		personService := service.NewPersonService(ageRepo, genderRepo, nationalityRepo, personRepo)
		personHandler := app.NewPersonHandler(ctx, personService, logger)

		Convey("Register person handler", func() {
			router := httprouter.New()
			personHandler.Register(router)

			Convey("Create person api method", func() {
				input := []byte(`
					{
						"name": "John",
						"surname": "Doe",
						"Patronymic": "John"
					}
				`)

				req := httptest.NewRequest(http.MethodPost, app.URL, bytes.NewBuffer(input))
				w := httptest.NewRecorder()

				personHandler.CreatePerson(w, req)

				So(w.Code, ShouldEqual, http.StatusCreated)

				var personResponse dto.PersonDTO

				_ = json.NewDecoder(w.Body).Decode(&personResponse)

				Convey("Update person api method", func() {
					input := []byte(`
						{
							"name": "John",
							"surname": "Doe",
							"patronymic": "John",
							"age": 32,
							"gender": "male",
							"nationality": "CA"
						}
					`)

					personID := personResponse.ID

					req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/persons/%s", personID), bytes.NewBuffer(input))
					w := httptest.NewRecorder()

					idParam := httprouter.Param{Key: "id", Value: personID}

					params := make([]httprouter.Param, 0)
					params = append(params, idParam)

					personHandler.UpdatePerson(w, req, httprouter.Params(params))

					So(w.Code, ShouldEqual, http.StatusOK)

					Convey("Get persons matches all filters", func() {
						req = httptest.NewRequest(http.MethodPost, "/api/v1/persons?age=32&gender=male&nationality=CA", nil)
						w = httptest.NewRecorder()

						personHandler.GetPersons(w, req)

						So(w.Code, ShouldEqual, http.StatusOK)
					})

					Convey("Get persons not matches age filters", func() {
						req = httptest.NewRequest(http.MethodPost, "/api/v1/persons?age=2323&gender=male&nationality=CA", nil)
						w = httptest.NewRecorder()

						personHandler.GetPersons(w, req)

						So(w.Code, ShouldEqual, http.StatusOK)
					})
				})

				Convey("Update empty person", func() {
					input := []byte("")

					personID := personResponse.ID

					req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/persons/%s", personID), bytes.NewBuffer(input))
					w := httptest.NewRecorder()

					idParam := httprouter.Param{Key: "id", Value: personID}

					params := make([]httprouter.Param, 0)
					params = append(params, idParam)

					personHandler.UpdatePerson(w, req, httprouter.Params(params))

					So(w.Code, ShouldEqual, http.StatusBadRequest)
				})

				Convey("Update person with empty fields", func() {
					input := []byte(`
						{
							"name": "",
							"surname": ""
						}
					`)

					personID := personResponse.ID

					req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/persons/%s", personID), bytes.NewBuffer(input))
					w := httptest.NewRecorder()

					idParam := httprouter.Param{Key: "id", Value: personID}

					params := make([]httprouter.Param, 0)
					params = append(params, idParam)

					personHandler.UpdatePerson(w, req, httprouter.Params(params))

					So(w.Code, ShouldEqual, http.StatusBadRequest)
				})

				Convey("Find person api method", func() {
					personID := personResponse.ID

					req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/persons/%s", personID), nil)
					w := httptest.NewRecorder()

					idParam := httprouter.Param{Key: "id", Value: personID}

					params := make([]httprouter.Param, 0)
					params = append(params, idParam)

					personHandler.FindPerson(w, req, httprouter.Params(params))

					So(w.Code, ShouldEqual, http.StatusOK)

					Convey("Delete person api method", func() {
						req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/persons/%s", personID), nil)
						w = httptest.NewRecorder()

						personHandler.DeletePerson(w, req, httprouter.Params(params))

						So(w.Code, ShouldEqual, http.StatusNoContent)
					})
				})
			})

			Convey("Update non-existent person", func() {
				input := []byte(`
						{
							"name": "John",
							"surname": "Doe",
							"patronymic": "John",
							"age": 32
						}
					`)

				personID := "asd"

				req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/persons/%s", personID), bytes.NewBuffer(input))
				w := httptest.NewRecorder()

				idParam := httprouter.Param{Key: "id", Value: personID}

				params := make([]httprouter.Param, 0)
				params = append(params, idParam)

				personHandler.UpdatePerson(w, req, httprouter.Params(params))

				So(w.Code, ShouldEqual, http.StatusBadRequest)
			})

			Convey("Find non-existent person", func() {
				personID := "asd"

				req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/persons/%s", personID), nil)
				w := httptest.NewRecorder()

				idParam := httprouter.Param{Key: "id", Value: personID}

				params := make([]httprouter.Param, 0)
				params = append(params, idParam)

				personHandler.FindPerson(w, req, httprouter.Params(params))

				So(w.Code, ShouldEqual, http.StatusBadRequest)

				Convey("Delete non-existent person", func() {

					req = httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/persons/%s", personID), nil)
					w = httptest.NewRecorder()

					personHandler.DeletePerson(w, req, httprouter.Params(params))

					So(w.Code, ShouldEqual, http.StatusBadRequest)
				})
			})

			Convey("Create empty person", func() {
				input := []byte("")

				req := httptest.NewRequest(http.MethodPost, app.URL, bytes.NewBuffer(input))
				w := httptest.NewRecorder()

				personHandler.CreatePerson(w, req)

				So(w.Code, ShouldEqual, http.StatusBadRequest)
			})

			Convey("Create person with empty fields", func() {
				input := []byte(`
					{
						"name": "",
						"surname": ""
					}
				`)

				req := httptest.NewRequest(http.MethodPost, app.URL, bytes.NewBuffer(input))
				w := httptest.NewRecorder()

				personHandler.CreatePerson(w, req)

				So(w.Code, ShouldEqual, http.StatusBadRequest)
			})
		})
	})
}
