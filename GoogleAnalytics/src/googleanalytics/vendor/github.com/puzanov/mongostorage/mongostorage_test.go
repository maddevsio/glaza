package mongostorage

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveStorage(t *testing.T) {
	s := NewItem()
	s.Host = "localhost"
	s.DB = "test"
	s.Name = "name"
	s.Value = "value"
	s.Collection = "docs"
	err := s.Save()
	assert.NoError(t, err)
}

func TestSaveStorageNestedJson(t *testing.T) {
	s := NewItem()
	s.Host = "localhost"
	s.DB = "test"
	s.Name = "name"
	s.JSON = "[{\"a\":\"b\"},{\"c\":\"d\"}]"
	s.Collection = "docs"
	err := s.Save()
	assert.NoError(t, err)
}

func TestSaveStorageBigNestedJson(t *testing.T) {
	var jsonData string
	dat, err := ioutil.ReadFile("test.json")
	assert.NoError(t, err)
	jsonData = string(dat)
	s := NewItem()
	s.Host = "localhost"
	s.DB = "test"
	s.Name = "name"
	s.JSON = jsonData
	s.Collection = "docs"
	err = s.Save()
	assert.NoError(t, err)
}
