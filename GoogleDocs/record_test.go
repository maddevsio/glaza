package main

import (
	"testing"
	"log"
	"github.com/stretchr/testify/assert"
)

func TestInvokeRecord(t *testing.T) {
	record := NewRecord()
	record.Doc = "https://drive.google.com/doc/sdfsdfsdf"
	record.Tab = "my tab"
	assert.Equal(t, "col1 | col2 | col3", record.getLastRecordInTab())
}

type Record struct {
	Doc string
	Tab string
}

func NewRecord() Record {
	return Record{}
}

func (r *Record) getLastRecordInTab() string {
	log.Print("implement last record")
	return "col1 | col2 | col3"
}