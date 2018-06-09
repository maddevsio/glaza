package main

import (
	"googlesheets/lib"
	"log"
	"os"
	"strings"
)

func main() {
	tab := lib.NewTab()
	tab.SpreadsheetID = os.Getenv("GOOGLE_SHEET_ID")
	tab.SecretFile = os.Getenv("SECRET_FILE")
	tab.Name = os.Getenv("TAB_NAME")

	items, err := tab.GetLastRecord()
	if err != nil {
		log.Fatal(err)
	}

	storage := lib.NewStorageItem()
	storage.Name = tab.Name
	storage.Collection = "docs"
	storage.Value = strings.Join(items, " ")
	err = storage.Save()
	if err != nil {
		log.Fatal(err)
	}
}
