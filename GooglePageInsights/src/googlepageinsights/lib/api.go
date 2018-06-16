package lib

import (
	"fmt"
	"log"
	"strconv"

	"github.com/buger/jsonparser"
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
// returns speed, json and error
func (a *API) GetPageInsights(url string, strategy string) (string, string, error) {
	payload := fmt.Sprintf(
		"https://www.googleapis.com/pagespeedonline/v4/runPagespeed?url=%v&strategy=%v&key=%v",
		url,
		strategy,
		a.Key,
	)
	resp, err := resty.R().Get(payload)
	if err != nil {
		log.Print("resty")
		return "", "", err
	}

	respBytes := resp.Body()
	speed, err := jsonparser.GetInt(respBytes, "ruleGroups", "SPEED", "score")
	if err != nil {
		log.Print("json parser")
		return "", "", err
	}

	return strconv.Itoa(int(speed)), string(respBytes), nil
}

/*
{
    "captchaResult": "CAPTCHA_NOT_NEEDED",
    "kind": "pagespeedonline#result",
    "id": "https://showmebishkek.com/",
    "responseCode": 200,
    "title": "Bishkek city tours, visit Ala-Archa gorge and Burana tower",
    "ruleGroups": {
     "SPEED": {
      "score": 68
     }
*/
