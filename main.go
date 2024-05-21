package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"time"
)

func main() {
	var records []Record
	readJson(&records)

	fileName := "parsed.xlsx"
	if !fileExists(fileName) {
		initExcel(fileName)
	}

	for _, j := range records {
		var price int
		currentTime := time.Now()
		formattedTime := currentTime.Format("15:04:05 02-01-2006")

		page := getPage(j.Url)

		data := findData(page, j.Tag, j.Substring)
		if data != nil {
			price = slices.Min(data)
		} else {
			fmt.Println("Empty")
		}

		er := excelRecord{
			Name:  j.Name,
			Price: price,
			Time:  formattedTime,
		}

		writeToExcel(fileName, er)

		fmt.Print("Press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
