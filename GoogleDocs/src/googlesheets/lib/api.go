package lib

import (
	"gopkg.in/Iwark/spreadsheet.v2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"context"
)

type API struct {
	SecretFile string
	SpreadsheetID string
}

func NewAPI() API {
	return API{}
}

func (a *API) GetDocument() (spreadsheet.Spreadsheet, error) {
	service, err := a.getSheetService()
	if err != nil {
		return spreadsheet.Spreadsheet{}, err
	}

	document, err := service.FetchSpreadsheet(a.SpreadsheetID)
	if err != nil {
		return spreadsheet.Spreadsheet{}, err
	}

	return document, nil
}

func (a *API) getSheetService() (*spreadsheet.Service, error) {
	data, err := ioutil.ReadFile(a.SecretFile)
	if err != nil {
		return nil, err
	}

	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	if err != nil {
		return nil, err
	}

	client := conf.Client(context.TODO())
	return spreadsheet.NewServiceWithClient(client), nil
}