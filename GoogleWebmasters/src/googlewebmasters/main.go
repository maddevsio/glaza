package main

import (
	"googlewebmasters/lib"
	"log"
	"os"

	"github.com/puzanov/mongostorage"
)

func main() {
	api := lib.NewAPI()
	api.SecretFile = os.Getenv("SECRET_FILE")
	api.URL = os.Getenv("URL")

	jsonData, err := api.GetClicksForNDays(30)
	if err != nil {
		log.Panic(err)
	}
	storage1 := mongostorage.NewItem()
	storage1.Name = os.Getenv("URL")
	storage1.Collection = "webmasters"
	storage1.Value = "clicks"
	storage1.JSON = jsonData
	err = storage1.Save()
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err = api.GetQueriesForNDays(30)
	if err != nil {
		log.Panic(err)
	}
	storage2 := mongostorage.NewItem()
	storage2.Name = os.Getenv("URL")
	storage2.Collection = "webmasters"
	storage2.Value = "queries"
	storage2.JSON = jsonData
	err = storage2.Save()
	if err != nil {
		log.Fatal(err)
	}
}
