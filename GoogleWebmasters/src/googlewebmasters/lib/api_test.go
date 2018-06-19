package lib

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApi(t *testing.T) {
	api := NewAPI()
	api.SecretFile = "client_secret.json"
	api.URL = "https:silkroadexplore.com"
	jsonData, err := api.GetDataForNDays(0)
	assert.NoError(t, err)
	assert.True(t, json.Valid([]byte(jsonData)))
	t.Log(jsonData)
}
