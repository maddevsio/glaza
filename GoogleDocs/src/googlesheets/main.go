package main

import (
	"googlesheets/lib"
	"log"
	"os"
)

func main() {
	log.Print("Empty logic")
	tab := lib.NewTab()
	tab.SpreadsheetID = os.Getenv("GOOGLE_SHEET_ID")
	tab.SecretFile = os.Getenv("SECRET_FILE")
	tab.Name = os.Getenv("TAB_NAME")
	log.Print(tab.GetLastRecord())

	// save data to mongo
	// get several record from doc, for example 5 last items
	// create different service who can make diffs within json structures
}
