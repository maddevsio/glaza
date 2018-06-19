package lib

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/go-resty/resty"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2/google"
)

func TestApi(t *testing.T) {
	data, err := ioutil.ReadFile("client_secret.json")
	assert.NoError(t, err)

	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/webmasters.readonly")
	assert.NoError(t, err)

	client := conf.Client(context.TODO())

	rc := resty.NewWithClient(client)
	restyResp, err := rc.R().
		SetBody(map[string]interface{}{"startDate": "2018-06-01", "endDate": "2018-06-18", "dimensions": []string{"date"}}).
		Post("https://www.googleapis.com/webmasters/v3/sites/https:silkroadexplore.com%2F/searchAnalytics/query")
	assert.NoError(t, err)
	t.Log(restyResp.StatusCode())
	t.Log(restyResp.String())

	t.Fail()
}
