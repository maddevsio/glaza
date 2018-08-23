package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPI(t *testing.T) {
	api := NewAPI()
	api.SecretFile = "client_secret.json"
	api.URL = "https://showmebishkek.com/"
	api.ViewID = "ga:163539924"

	ga, err := api.GetDataFor7Days()
	assert.NoError(t, err)
	assert.Contains(t, ga.JSON, "\"totalResults\":7,")
}
