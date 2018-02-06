package lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetLastRecord(t *testing.T) {
	tab := NewTab()
	tab.SecretFile = "client_secret.json"
	tab.SpreadsheetID = "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
	tab.Name = "Class Data"
	lastRecord, err := tab.GetLastRecord()
	assert.NoError(t, err)
	assert.Len(t, lastRecord, 6)
	assert.Equal(t, "Will", lastRecord[0])
	assert.Equal(t, "Debate", lastRecord[len(lastRecord) - 1])
}