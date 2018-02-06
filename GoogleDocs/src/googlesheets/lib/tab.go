package lib

import "log"

type Tab struct {
	API
	Name string
}

func NewTab() Tab {
	return Tab{}
}

func (r *Tab) GetLastRecord() string {
	log.Print("implement last record")
	return "col1 | col2 | col3"
}
