package main

import (
	"fmt"
	"slices"
)

func main() {
	var records []Record

	readJson(&records)
	for _, j := range records {
		page := getPageFromFile(j.Url)
		data := findData(page, j.Tag, j.Substring)
		fmt.Println(slices.Min(data))
	}
}
