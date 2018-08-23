package lib

import (
	"context"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2/google"
	analytics "google.golang.org/api/analytics/v3"
)

//https://developers.google.com/analytics/devguides/reporting/core/dimsmets#cats=user,time,session

// GA is a root struct for Google Analytics data
type GA struct {
	URL     string
	JSON    string
	GADatas []GAData
}

// GAData is a struct which holds particular Google Analytics data for a given view
type GAData struct {
	Date            string
	Users           string
	AdImpressions   string
	AdClicks        string
	OrganicSearches string
}

// API is a main service to query Google Analytics
type API struct {
	SecretFile string
	URL        string
	ViewID     string
}

// NewAPI is a API struct constructor
func NewAPI() API {
	return API{}
}

// GetDataFor7Days json from GoogleAnalytics with needed raw data
func (a *API) GetDataFor7Days() (GA, error) {
	client, err := a.GetOAuthClient()
	if err != nil {
		return GA{}, err
	}

	analyticsService, err := analytics.New(client)
	if err != nil {
		return GA{}, err
	}

	gaData := analyticsService.Data.Ga.Get(a.ViewID, "7daysAgo", "yesterday",
		"ga:users, ga:impressions, ga:adClicks, ga:organicSearches")
	gaData.Dimensions("ga:day")
	gaData.SamplingLevel("HIGHER_PRECISION")
	gaData.Output("json")

	d, err := gaData.Do()
	if err != nil {
		return GA{}, err
	}

	bytes, err := d.MarshalJSON()
	if err != nil {
		return GA{}, err
	}

	gaDatas := []GAData{}

	for _, value := range d.Rows {
		gaData := GAData{}
		gaData.Date = value[0]
		gaData.Users = value[1]
		gaData.AdImpressions = value[2]
		gaData.AdClicks = value[3]
		gaData.OrganicSearches = value[4]
		gaDatas = append(gaDatas, gaData)
	}

	ga := GA{}
	ga.GADatas = gaDatas
	ga.JSON = string(bytes)
	ga.URL = a.URL

	return ga, nil
}

// GetOAuthClient returns valid http.client ready to work with google api
func (a *API) GetOAuthClient() (*http.Client, error) {
	data, err := ioutil.ReadFile(a.SecretFile)
	if err != nil {
		return nil, err
	}

	conf, err := google.JWTConfigFromJSON(data, analytics.AnalyticsReadonlyScope)
	if err != nil {
		return nil, err
	}

	client := conf.Client(context.TODO())
	return client, nil
}
