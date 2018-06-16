package lib

import (
	"fmt"

	"github.com/go-resty/resty"
)

// API is used to make queries to a GooglePageInsights API
// here we need Key from Google API Console
type API struct {
	Key string
}

// NewAPI is a constructor for API
func NewAPI(key string) API {
	api := API{}
	api.Key = key
	return api
}

// GetPageInsights makes an API call to Google PageInsights server
func (a *API) GetPageInsights(url string, strategy string) (string, error) {
	payload := fmt.Sprintf(
		"https://www.googleapis.com/pagespeedonline/v4/runPagespeed?url=%v&strategy=%v&key=%v",
		url,
		strategy,
		a.Key,
	)
	resp, err := resty.R().Get(payload)
	if err != nil {
		return "", nil
	}
	return resp.String(), nil
}
