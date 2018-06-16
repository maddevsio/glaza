package lib

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	assert.NoError(t, err)

	key := viper.GetString("key")
	api := NewAPI(key)
	speed, json, err := api.GetPageInsights("http://showmebishkek.com", "desktop")
	assert.NoError(t, err)
	assert.Equal(t, "68", speed)
	assert.Contains(t, json, "\"responseCode\": 200,")
	assert.Contains(t, json, "\"id\": \"https://showmebishkek.com/\",")
}
