package lib

type Tab struct {
	API
	Name string
}

func NewTab() Tab {
	return Tab{}
}

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
	lastRow := sheet.Rows[len(sheet.Rows) - 1]
	for _, cell := range lastRow {
		if cell.Value != "" {
			lastRecord = append(lastRecord, cell.Value)
		}
	}
	return lastRecord, nil
}