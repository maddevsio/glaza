package lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit#gid=0
// this is example doc from Google for Google Sheets API testing
func getAPI() API {
	api := NewAPI()
	api.SpreadsheetID = "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
	api.SecretFile = "client_secret.json"
	return api
}

func TestRunApi_GetTabByIndex(t *testing.T) {
	api := getAPI()
	doc, err := api.GetDocument()
	assert.NoError(t, err)
	sheet, err := doc.SheetByIndex(0)
	assert.NoError(t, err)
	assert.Equal(t, "Alexandra", sheet.Rows[1][0].Value)
}