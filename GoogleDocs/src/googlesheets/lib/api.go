package lib

import (
	"context"
	"io/ioutil"

	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

// API is a struct for Google Spreadsheet API
type API struct {
	SecretFile    string
	SpreadsheetID string
}

// NewAPI is a API struct constructor
func NewAPI() API {
	return API{}
}

// GetDocument returns a spreadsheet struct
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

// getSheetService is a private func for GetDocument. API.SecretFile filed should not be blank
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
