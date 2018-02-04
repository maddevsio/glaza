package lib

import "log"

type Tab struct {
	GoogleSheetAddress string
	Name               string
}

func NewTab() Tab {
	return Tab{}
}

func (r *Tab) GetLastRecord() string {
	log.Print("implement last record")
	return "col1 | col2 | col3"
}
