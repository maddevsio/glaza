package lib

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClicksForNDays(t *testing.T) {
	api := NewAPI()
	api.SecretFile = "client_secret.json"
	api.URL = "https:silkroadexplore.com"
	jsonData, err := api.GetClicksForNDays(30)
	assert.NoError(t, err)
	assert.True(t, json.Valid([]byte(jsonData)))
	assert.Contains(t, jsonData, "impressions")
	assert.Contains(t, jsonData, "clicks")
	t.Log(jsonData)
}

func TestGetQueriesForNDays(t *testing.T) {
	api := NewAPI()
	api.SecretFile = "client_secret.json"
	api.URL = "https:silkroadexplore.com"
	jsonData, err := api.GetQueriesForNDays(30)
	assert.NoError(t, err)
	assert.True(t, json.Valid([]byte(jsonData)))
	assert.Contains(t, jsonData, "keys")
	assert.Contains(t, jsonData, "impressions")
	assert.Contains(t, jsonData, "clicks")
	t.Log(jsonData)
}

func TestDate(t *testing.T) {
	tf := GetCurrentDate()
	assert.Contains(t, tf, "2018")
	tf = GetDateNDaysBefore(30)
	assert.Contains(t, tf, "2018") // lame
	t.Log(tf)
}
