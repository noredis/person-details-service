package person_app

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	service "person-details-service/internal/service/person"
	dto "person-details-service/internal/service/person/dto"
	"time"

	"github.com/julienschmidt/httprouter"
)

const URL = "/api/v1/persons"

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
}

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
