package lib

// Tab is a struct for working with tabs in Google Spreadsheets. It references an API struct
type Tab struct {
	API
	Name string
}

// NewTab is a Tab struct constructor
func NewTab() Tab {
	return Tab{}
}

// GetLastRecord returns the last record of a tab in a diven document. Document ID should be passed via API struct
func (tab *Tab) GetLastRecord() ([]string, error) {
	doc, err := tab.GetDocument()
	if err != nil {
		return nil, err
	}

	sheet, err := doc.SheetByTitle(tab.Name)
	if err != nil {
		return nil, err
	}

	var lastRecord []string
	lastRow := sheet.Rows[len(sheet.Rows)-1]
	for _, cell := range lastRow {
		if cell.Value != "" {
			lastRecord = append(lastRecord, cell.Value)
		}
	}
	return lastRecord, nil
}
