package main

import (
	"googleanalytics/lib"
	"log"
	"os"

	"github.com/puzanov/mongostorage"
)

func main() {
	api := lib.NewAPI()
	api.SecretFile = os.Getenv("SECRET_FILE")
	api.URL = os.Getenv("URL")
	api.ViewID = os.Getenv("VIEW_ID")

	ga, err := api.GetDataFor7Days()
	if err != nil {
		log.Fatal(err)
	}

	storage := mongostorage.NewItem()
	storage.Name = ga.URL
	storage.Collection = "analytics"
	storage.Value = api.ViewID
	storage.JSON = string(ga.JSON)
	err = storage.Save()
	if err != nil {
		log.Fatal(err)
	}
}
