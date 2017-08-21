package output

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func toXls(dest string, items []map[string]interface{}) error {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	savePath, err := getSavePath(dest)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		fmt.Println("nothing")
		return nil
	}
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		return err
	}
	firstRow := items[0]
	row = sheet.AddRow()
	var xlsColumns []string
	for colName := range firstRow {
		cell = row.AddCell()
		cell.Value = colName
		xlsColumns = append(xlsColumns, colName)
	}

	for _, item := range items {
		row = sheet.AddRow()
		for i := 0; i < len(xlsColumns); i++ {
			val := item[xlsColumns[i]]
			cell = row.AddCell()
			cell.Value = val.(string)
		}
	}
	fmt.Println("Saved to ", savePath)
	return file.Save(savePath)
}
