package main

import (
	"testing"
	"log"
	"github.com/stretchr/testify/assert"
)

func TestSmoke(t *testing.T) {
	tab := NewTab()
	tab.GoogleDocAddress = "https://drive.google.com/doc/sdfsdfsdf"
	tab.Name = "my tab"
	assert.Equal(t, "col1 | col2 | col3", tab.getLastRecord())
}

type Tab struct {
	GoogleDocAddress string
	Name             string
}

func NewTab() Tab {
	return Tab{}
}

func (r *Tab) getLastRecord() string {
	log.Print("implement last record")
	return "col1 | col2 | col3"
}