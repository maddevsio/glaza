package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveStorage(t *testing.T) {
	s := Storage{"line", "value"}
	err := s.Save()
	assert.NoError(t, err)
}
