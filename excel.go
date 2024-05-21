package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"strconv"
)

type excelRecord struct {
	Name  string
	Price int
	Time  string
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func initExcel(fileName string) *excelize.File {
	if !fileExists(fileName) {
		xf := excelize.NewFile()
		defer func() {
			if err := xf.Close(); err != nil {
				log.Println(err)
			}
		}()

		index, err := xf.NewSheet("Sheet1")
		if err != nil {
			log.Println(err)
		}

		xf.SetCellValue("Sheet1", "A1", "Название")
		xf.SetCellValue("Sheet1", "B1", "Цена")
		xf.SetCellValue("Sheet1", "C1", "Дата проверки")

		xf.SetActiveSheet(index)

		if err := xf.SaveAs(fileName); err != nil {
			log.Println(err)
		}
		return xf
	}
	return nil
}

func writeToExcel(fileName string, er excelRecord) {
	var sheetName = "Sheet1"
	xf, err := excelize.OpenFile(fileName)
	if err != nil {
		log.Println(err)
	}

	defer func() {
		if err = xf.Close(); err != nil {
			log.Println(err)
		}
	}()

	emptyRow, err := findFirstEmptyRow(xf, sheetName)
	if err != nil {
		fmt.Println(err)
		return
	}

	rowNum := strconv.Itoa(emptyRow)
	xf.SetCellValue(sheetName, "A"+rowNum, er.Name)
	xf.SetCellValue(sheetName, "B"+rowNum, er.Price)
	xf.SetCellValue(sheetName, "C"+rowNum, er.Time)

	if err = xf.SaveAs(fileName); err != nil {
		log.Println(err)
	}
}

func isRowEmpty(row []string) bool {
	for _, cell := range row {
		if cell != "" {
			return false
		}
	}
	return true
}

func findFirstEmptyRow(file *excelize.File, sheetName string) (int, error) {
	rows, err := file.GetRows(sheetName)
	if err != nil {
		return 0, err
	}

	for i, row := range rows {
		if isRowEmpty(row) {
			return i + 1, nil
		}
	}

	return len(rows) + 1, nil
}
