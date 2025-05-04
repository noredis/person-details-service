package gender_infra

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"person-details-service/internal/domain/person/valueobject"
	"time"
)

type genderResponse struct {
	Gender *string `json:"gender"`
}

type GenderRepository struct {
	client  *http.Client
	logger  slog.Logger
	baseURL *url.URL
}

func NewGenderRepository(logger slog.Logger, baseURL string, timeout time.Duration) (*GenderRepository, error) {
	const op = "gender_infra.NewGenderRepository: %w"

	reqURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	return &GenderRepository{
		client: &http.Client{
			Timeout: timeout,
		},
		logger:  logger,
		baseURL: reqURL,
	}, nil
}

func (r *GenderRepository) FindOutPersonsGender(ctx context.Context, fullName person_vo.FullName) (*person_vo.Gender, error) {
	const op = "GenderRepository.FindOutPersonsGender"

	reqURL := *r.baseURL

	query := reqURL.Query()
	query.Add("name", fullName.Value())
	reqURL.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to create request: %w", op, err)
	}

	r.logger.Debug("Sending request to gender API",
		"url", reqURL.String(),
		"person", fullName.Value(),
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
		r.logger.Error("Unable to read gender API response body",
			"status", resp.StatusCode,
			"url", reqURL.String(),
		)
	}

	if resp.StatusCode != http.StatusOK {
		r.logger.Error("Gender API request completed with error",
			"status", resp.StatusCode,
			"url", reqURL.String(),
			"response", string(body),
		)

		return nil, fmt.Errorf("%s: API returned status code: %d", op, resp.StatusCode)
	}

	r.logger.Debug("Gender API request completed successfully",
		"url", reqURL.String(),
		"response", string(body),
	)

	var genderResp genderResponse

	if err := json.Unmarshal(body, &genderResp); err != nil {
		return nil, fmt.Errorf("%s: unable to decode response body: %w", op, err)
	}

	if genderResp.Gender == nil {
		return nil, nil
	}

	gender, err := person_vo.NewGender(*genderResp.Gender)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return gender, nil
}
