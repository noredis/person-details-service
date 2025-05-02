package nationality_infra

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

type country struct {
	CountryID string `json:"country_id"`
}

type nationalityResponse struct {
	Country *[]country `json:"country"`
}

type NationalityRepository struct {
	client  *http.Client
	logger  slog.Logger
	baseURL *url.URL
}

func NewNationalityRepository(logger slog.Logger, baseURL string, timeout time.Duration) (*NationalityRepository, error) {
	const op = "nationality_infra.NewNationalityRepository: %w"

	reqURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	return &NationalityRepository{
		client: &http.Client{
			Timeout: timeout,
		},
		logger:  logger,
		baseURL: reqURL,
	}, nil
}

func (r *NationalityRepository) FindOutPersonsNationality(ctx context.Context, fullName person_vo.FullName) (*person_vo.Nationality, error) {
	const op = "NationalityRepository.FindOutPersonsNationality"

	reqURL := *r.baseURL

	query := reqURL.Query()
	query.Add("name", fullName.Value())
	reqURL.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to create request: %w", op, err)
	}

	r.logger.Debug("Sending request to nationality API",
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
		r.logger.Error("Unable to read nationality API response body",
			"status", resp.StatusCode,
			"url", reqURL.String(),
		)
	}

	if resp.StatusCode != http.StatusOK {
		r.logger.Error("Nationality API request completed with error",
			"status", resp.StatusCode,
			"url", reqURL.String(),
			"response", string(body),
		)

		return nil, fmt.Errorf("%s: API returned status code: %d", op, resp.StatusCode)
	}

	r.logger.Debug("Nationality API request completed successfully",
		"url", reqURL.String(),
		"response", string(body),
	)

	var nationalityResp nationalityResponse

	if err := json.Unmarshal(body, &nationalityResp); err != nil {
		return nil, fmt.Errorf("%s: unable to decode response body: %w", op, err)
	}

	if nationalityResp.Country == nil || len(*nationalityResp.Country) == 0 {
		return nil, nil
	}

	nationality, err := person_vo.NewNationality((*nationalityResp.Country)[0].CountryID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return nationality, nil
}
