package person_app

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	repo "person-details-service/internal/repo/person"
	service "person-details-service/internal/service/person"
	dto "person-details-service/internal/service/person/dto"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

const URL = "/api/v1/persons"
const URL_ID = "/api/v1/persons/:id"

type PersonHandler struct {
	ctx           context.Context
	personService service.PersonService
	logger        *slog.Logger
}

func NewPersonHandler(ctx context.Context, personService service.PersonService, logger *slog.Logger) PersonHandler {
	return PersonHandler{ctx, personService, logger}
}

func (h PersonHandler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, URL, h.CreatePerson)
	router.PUT(URL_ID, h.UpdatePerson)
	router.GET(URL_ID, h.FindPerson)
	router.DELETE(URL_ID, h.DeletePerson)
	router.HandlerFunc(http.MethodGet, URL, h.GetPersons)
}

// CreatePerson godoc
// @Summary Create a new person
// @Description Create a new person with the input payload
// @Tags persons
// @Accept json
// @Produce json
// @Param X-Request-ID header string false "Request ID"
// @Param person body dto.CreatePersonDTO true "Person data to create"
// @Success 201 {object} dto.PersonDTO
// @Header 201 {string} Location "URL to the created resource"
// @Failure 400 {string} string "Bad request"
// @Router /persons [post]
func (h PersonHandler) CreatePerson(w http.ResponseWriter, req *http.Request) {
	requestID := req.Header.Get("X-Request-ID")
	startTime := time.Now()

	logger := h.logger.With(
		slog.String("request_id", requestID),
		slog.String("method", req.Method),
		slog.String("path", req.URL.Path),
		slog.String("remote_ip", req.RemoteAddr),
	)

	logger.Info("Request to create Person has been received")

	var createPersonDTO dto.CreatePersonDTO
	err := json.NewDecoder(req.Body).Decode(&createPersonDTO)
	if err != nil {
		logger.Error("Unable to decode request body",
			slog.String("error", err.Error()))

		w.WriteHeader(http.StatusBadRequest)

		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			logger.Error("Unable to write request response",
				slog.String("error", err.Error()))
		}
		return
	}

	logger.Debug("Request to create Person has been decoded",
		slog.Any("dto", createPersonDTO))

	personDTO, err := h.personService.CreatePerson(h.ctx, createPersonDTO)
	if err != nil {
		logger.Error("Unable to create Person",
			slog.String("error", err.Error()),
			slog.Any("dto", createPersonDTO))

		w.WriteHeader(http.StatusBadRequest)

		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			logger.Error("Unable to write request response",
				slog.String("error", err.Error()))
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", fmt.Sprintf("%s/%s", URL, personDTO.ID))
	w.WriteHeader(http.StatusCreated)

	encodeErr := json.NewEncoder(w).Encode(personDTO)
	if encodeErr != nil {
		logger.Error("Unable to encode request response",
			slog.String("error", err.Error()))
		return
	}

	logger.Info("Person has been created successfully",
		slog.String("person_id", personDTO.ID),
		slog.Int64("duration_ms", time.Since(startTime).Milliseconds()),
		slog.Int("status_code", http.StatusCreated))
}

// UpdatePerson godoc
// @Summary Update an existing person
// @Description Update person details by ID
// @Tags persons
// @Accept json
// @Produce json
// @Param X-Request-ID header string false "Request ID"
// @Param id path string true "Person ID"
// @Param person body dto.UpdatePersonDTO true "Person data to update"
// @Success 200 {object} dto.PersonDTO
// @Failure 400 {string} string "Bad request"
// @Router /persons/{id} [put]
func (h PersonHandler) UpdatePerson(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	requestID := req.Header.Get("X-Request-ID")
	startTime := time.Now()

	logger := h.logger.With(
		slog.String("request_id", requestID),
		slog.String("method", req.Method),
		slog.String("path", req.URL.Path),
		slog.String("remote_ip", req.RemoteAddr),
	)

	logger.Info("Request to update Person has been received")

	var updatePersonDTO dto.UpdatePersonDTO

	err := json.NewDecoder(req.Body).Decode(&updatePersonDTO)
	if err != nil {
		logger.Error("Unable to decode request body",
			slog.String("error", err.Error()))

		w.WriteHeader(http.StatusBadRequest)

		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			logger.Error("Unable to write request response",
				slog.String("error", err.Error()))
		}
		return
	}

	logger.Debug("Request to update Person has been decoded",
		slog.Any("dto", updatePersonDTO))

	personDTO, err := h.personService.UpdatePerson(h.ctx, ps.ByName("id"), updatePersonDTO)
	if err != nil {
		logger.Error("Unable to update Person",
			slog.String("error", err.Error()),
			slog.Any("dto", updatePersonDTO))

		w.WriteHeader(http.StatusBadRequest)

		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			logger.Error("Unable to write request response",
				slog.String("error", err.Error()))
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encodeErr := json.NewEncoder(w).Encode(personDTO)
	if encodeErr != nil {
		logger.Error("Unable to encode request response",
			slog.String("error", err.Error()))
		return
	}

	logger.Info("Person has been updated successfully",
		slog.String("person_id", personDTO.ID),
		slog.Int64("duration_ms", time.Since(startTime).Milliseconds()),
		slog.Int("status_code", http.StatusCreated))
}

// FindPerson godoc
// @Summary Get person by ID
// @Description Get person details by ID
// @Tags persons
// @Produce json
// @Param X-Request-ID header string false "Request ID"
// @Param id path string true "Person ID"
// @Success 200 {object} dto.PersonDTO
// @Failure 400 {string} string "Bad request"
// @Router /persons/{id} [get]
func (h PersonHandler) FindPerson(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	requestID := req.Header.Get("X-Request-ID")
	startTime := time.Now()

	logger := h.logger.With(
		slog.String("request_id", requestID),
		slog.String("method", req.Method),
		slog.String("path", req.URL.Path),
		slog.String("remote_ip", req.RemoteAddr),
	)

	logger.Info("Request to find Person has been received")

	personDTO, err := h.personService.FindPerson(h.ctx, ps.ByName("id"))
	if err != nil {
		logger.Error("Unable to find Person",
			slog.String("error", err.Error()))

		w.WriteHeader(http.StatusBadRequest)

		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			logger.Error("Unable to write request response",
				slog.String("error", err.Error()))
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encodeErr := json.NewEncoder(w).Encode(personDTO)
	if encodeErr != nil {
		logger.Error("Unable to encode request response",
			slog.String("error", err.Error()))
		return
	}

	logger.Info("Person has been finded successfully",
		slog.String("person_id", personDTO.ID),
		slog.Int64("duration_ms", time.Since(startTime).Milliseconds()),
		slog.Int("status_code", http.StatusCreated))
}

// DeletePerson godoc
// @Summary Delete person by ID
// @Description Delete person by ID
// @Tags persons
// @Param X-Request-ID header string false "Request ID"
// @Param id path string true "Person ID"
// @Success 204 "No Content"
// @Failure 400 {string} string "Bad request"
// @Router /persons/{id} [delete]
func (h PersonHandler) DeletePerson(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	requestID := req.Header.Get("X-Request-ID")
	startTime := time.Now()

	logger := h.logger.With(
		slog.String("request_id", requestID),
		slog.String("method", req.Method),
		slog.String("path", req.URL.Path),
		slog.String("remote_ip", req.RemoteAddr),
	)

	logger.Info("Request to delete Person has been received")

	err := h.personService.DeletePerson(h.ctx, ps.ByName("id"))
	if err != nil {
		logger.Error("Unable to delete Person",
			slog.String("error", err.Error()))

		w.WriteHeader(http.StatusBadRequest)

		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			logger.Error("Unable to write request response",
				slog.String("error", err.Error()))
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

	logger.Info("Person has been deleted successfully",
		slog.String("person_id", ps.ByName("id")),
		slog.Int64("duration_ms", time.Since(startTime).Milliseconds()),
		slog.Int("status_code", http.StatusCreated))
}

// GetPersons godoc
// @Summary Get list of persons
// @Description Get list of persons with optional filtering
// @Tags persons
// @Produce json
// @Param X-Request-ID header string false "Request ID"
// @Param age query integer false "Filter by age"
// @Param gender query string false "Filter by gender"
// @Param nationality query string false "Filter by nationality"
// @Success 200 {array} dto.PersonDTO
// @Failure 400 {string} string "Bad request"
// @Router /persons [get]
func (h *PersonHandler) GetPersons(w http.ResponseWriter, req *http.Request) {
	requestID := req.Header.Get("X-Request-ID")
	startTime := time.Now()

	logger := h.logger.With(
		slog.String("request_id", requestID),
		slog.String("method", req.Method),
		slog.String("path", req.URL.Path),
		slog.String("remote_ip", req.RemoteAddr),
	)

	logger.Info("Request to find Persons has been received")

	query := req.URL.Query()

	var filters repo.FilterOptions

	if ageStr := query.Get("age"); ageStr != "" {
		age, err := strconv.Atoi(ageStr)
		if err == nil {
			filters.Age = &age
		}
	}

	if gender := query.Get("gender"); gender != "" {
		filters.Gender = &gender
	}

	if nationality := query.Get("nationality"); nationality != "" {
		filters.Nationality = &nationality
	}

	personDTOs, err := h.personService.GetPersons(h.ctx, filters)
	if err != nil {
		logger.Error("Unable to find Persons",
			slog.String("error", err.Error()))

		w.WriteHeader(http.StatusBadRequest)

		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			logger.Error("Unable to write request response",
				slog.String("error", err.Error()))
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encodeErr := json.NewEncoder(w).Encode(personDTOs)
	if encodeErr != nil {
		logger.Error("Unable to encode request response",
			slog.String("error", err.Error()))
		return
	}

	logger.Info("Persons has been finded successfully",
		slog.Int64("duration_ms", time.Since(startTime).Milliseconds()),
		slog.Int("status_code", http.StatusCreated))
}
