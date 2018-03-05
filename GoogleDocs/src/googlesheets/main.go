package main

import (
	"googlesheets/lib"
	"log"
)

func main() {
	log.Print("Empty logic")
	tab := lib.NewTab()
	log.Print(tab.GetLastRecord())
}
