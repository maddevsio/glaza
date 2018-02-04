package main

import (
	"log"
	"googlesheets/lib"
)

func main() {
	log.Print("Empty logic")
	tab := lib.NewTab()
	log.Print(tab.GetLastRecord())
}
