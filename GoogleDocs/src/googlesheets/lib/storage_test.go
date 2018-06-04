package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveStorage(t *testing.T) {
	s := NewStorageItem()
	s.Name = "name"
	s.Value = "value"
	s.Collection = "docs"
	err := s.Save()
	assert.NoError(t, err)
}
