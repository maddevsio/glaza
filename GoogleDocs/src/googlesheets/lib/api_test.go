package lib

import (
	"golang.org/x/oauth2/google"
	"testing"
	"context"
	"log"
	"io/ioutil"
	"gopkg.in/Iwark/spreadsheet.v2"
)


func TestRunApi(t *testing.T) {
	data, err := ioutil.ReadFile("client_secret.json")
	checkError(err)

	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	checkError(err)

	client := conf.Client(context.TODO())
	service := spreadsheet.NewServiceWithClient(client)

	spreadsheetID := "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
	spreadsheet, err := service.FetchSpreadsheet(spreadsheetID)

	sheet, err := spreadsheet.SheetByIndex(0)

	log.Printf("cell value: %v", sheet.Rows[0][1].Value)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}