package lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSmoke(t *testing.T) {
	tab := NewTab()
	tab.GoogleSheetAddress = "https://drive.google.com/doc/sdfsdfsdf"
	tab.Name = "my tab"
	assert.Equal(t, "col1 | col2 | col3", tab.GetLastRecord())
}