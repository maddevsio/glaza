package lib

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

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

// GetClicksForNDays returns search analytics for a number of days passed as a param
// https://developers.google.com/webmaster-tools/search-console-api-original/v3/how-tos/search_analytics
func (a *API) GetClicksForNDays(days int) (string, error) {
	client, err := a.GetOAuthClient()
	if err != nil {
		return "", err
	}
	rc := resty.NewWithClient(client)
	restyResp, err := rc.R().
		SetBody(map[string]interface{}{"startDate": GetDateNDaysBefore(days), "endDate": GetCurrentDate(), "dimensions": []string{"date"}}).
		Post("https://www.googleapis.com/webmasters/v3/sites/" + a.URL + "%2F/searchAnalytics/query")

	return restyResp.String(), err
}

// GetQueriesForNDays returns search analytics for a number of days passed as a param
func (a *API) GetQueriesForNDays(days int) (string, error) {
	client, err := a.GetOAuthClient()
	if err != nil {
		return "", err
	}
	rc := resty.NewWithClient(client)
	restyResp, err := rc.R().
		SetBody(map[string]interface{}{"startDate": GetDateNDaysBefore(days), "endDate": GetCurrentDate(), "dimensions": []string{"query"}}).
		Post("https://www.googleapis.com/webmasters/v3/sites/" + a.URL + "%2F/searchAnalytics/query")

	return restyResp.String(), err
}

// GetOAuthClient returns valid http.client ready to work with google api
func (a *API) GetOAuthClient() (*http.Client, error) {
	data, err := ioutil.ReadFile(a.SecretFile)
	if err != nil {
		return nil, err
	}

	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/webmasters.readonly")
	if err != nil {
		return nil, err
	}

	client := conf.Client(context.TODO())
	return client, nil
}

// GetCurrentDate returns current date in format 2018-31-12
func GetCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

// GetDateNDaysBefore returns date in format 2018-31-12 minus N days from current date
func GetDateNDaysBefore(days int) string {
	return time.Now().AddDate(0, 0, -days).Format("2006-01-02")
}
