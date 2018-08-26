package main

import (
	"fmt"
	"googlesheets/lib"
	"log"
	"os"
	"strings"

	"github.com/puzanov/mongostorage"
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

	storage := mongostorage.NewItem()
	storage.Name = tab.Name
	storage.Collection = "docs"
	storage.Value = strings.Join(items, " ")
	storage.JSON = fmt.Sprintf("{\"url\":\"%v\"}", os.Getenv("URL"))
	err = storage.Save()
	if err != nil {
		log.Fatal(err)
	}
}
