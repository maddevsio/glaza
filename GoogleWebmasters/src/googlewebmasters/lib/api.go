package lib

import (
	"context"
	"io/ioutil"

	"github.com/go-resty/resty"
	"golang.org/x/oauth2/google"
)

// API is a struct for Google Webmasters API
// URL goes in this format: https:example.com,
// because of a bug in resty: https://github.com/go-resty/resty/issues/159
type API struct {
	SecretFile string
	URL        string
}

// NewAPI is a API struct constructor
func NewAPI() API {
	return API{}
}

// GetDataForNDays returns search analytics for a number of days passed as a param
func (a *API) GetDataForNDays(days int) (string, error) {
	data, err := ioutil.ReadFile(a.SecretFile)
	if err != nil {
		return "", err
	}

	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/webmasters.readonly")
	if err != nil {
		return "", err
	}

	client := conf.Client(context.TODO())
	rc := resty.NewWithClient(client)
	restyResp, err := rc.R().
		SetBody(map[string]interface{}{"startDate": "2018-06-01", "endDate": "2018-06-18", "dimensions": []string{"date"}}).
		Post("https://www.googleapis.com/webmasters/v3/sites/" + a.URL + "%2F/searchAnalytics/query")

	return restyResp.String(), err
}

// GetCurrentDate returns current date in format 2018-31-12
func GetCurrentDate() string {
	return "2018"
}

// GetDateNDaysBefore returns date in format 2018-31-12 minus N days from current date
func GetDateNDaysBefore(days int) string {
	return "2039"
}
