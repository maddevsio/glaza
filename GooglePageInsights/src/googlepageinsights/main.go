package main

import (
	"googlepageinsights/lib"
	"log"
	"os"

	"github.com/puzanov/mongostorage"
)

func main() {
	api := lib.NewAPI(os.Getenv("KEY"))
	speed, json, err := api.GetPageInsights(os.Getenv("URL"), os.Getenv("STRATEGY"))
	if err != nil {
		log.Fatal(err)
	}

	storage := mongostorage.NewItem()
	storage.Name = os.Getenv("URL") + " " + os.Getenv("STRATEGY")
	storage.Collection = "pageinsights"
	storage.Value = speed
	storage.JSON = json
	err = storage.Save()
	if err != nil {
		log.Fatal(err)
	}
}
