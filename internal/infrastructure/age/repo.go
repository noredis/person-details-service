package age_infra

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"person-details-service/internal/domain/person"
	"person-details-service/internal/domain/person/valueobject"
	"time"
)

type ageResponse struct {
	Age *int `json:"age"`
}

type AgeRepository struct {
	client  *http.Client
	logger  slog.Logger
	baseURL *url.URL
}

func NewAgeRepository(logger slog.Logger, baseURL string, timeout time.Duration) (*AgeRepository, error) {
	const op = "age_infra.NewAgeRepository: %w"

	reqURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	return &AgeRepository{
		client: &http.Client{
			Timeout: timeout,
		},
		logger:  logger,
		baseURL: reqURL,
	}, nil
}

func (r *AgeRepository) FindOutPersonsAge(ctx context.Context, p person.Person) (*person_vo.Age, error) {
	const op = "AgeRepository.FindOutPersonsAge"

	reqURL := *r.baseURL

	query := reqURL.Query()
	query.Add("name", p.FullName().Value())
	reqURL.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to create request: %w", op, err)
	}

	r.logger.Debug("Sending request to age API",
		"url", reqURL.String(),
		"person", p.FullName(),
	)

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to send request: %w", op, err)
	}

	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			r.logger.Warn(
				"Unable to close response body",
				"error", closeErr,
				"url", req.URL.String(),
				"operation", "HTTP response closing",
			)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		r.logger.Error("Unable to read age API response body",
			"status", resp.StatusCode,
			"url", reqURL.String(),
		)
	}

	if resp.StatusCode != http.StatusOK {
		r.logger.Error("Age API request completed with error",
			"status", resp.StatusCode,
			"url", reqURL.String(),
			"response", string(body),
		)

		return nil, fmt.Errorf("%s: API returned status code: %d", op, resp.StatusCode)
	}

	r.logger.Debug("Age API request completed successfully",
		"url", reqURL.String(),
		"response", string(body),
	)

	var ageResp ageResponse

	if err := json.Unmarshal(body, &ageResp); err != nil {
		return nil, fmt.Errorf("%s: unable to decode response body: %w", op, err)
	}

	if ageResp.Age == nil {
		return nil, nil
	}

	age, err := person_vo.NewAge(*ageResp.Age)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return age, nil
}
