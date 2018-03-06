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
	// TODO: see readme for things to be done here
}
