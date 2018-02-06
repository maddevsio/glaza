package lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSmoke(t *testing.T) {
	tab := NewTab()
	tab.SecretFile = "client_secret.json"
	tab.SpreadsheetID = "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
	tab.Name = "Class Data"
	assert.Equal(t, "col1 | col2 | col3", tab.GetLastRecord())
}