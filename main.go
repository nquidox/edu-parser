package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	var records []Record
	readJson(&records)
	for _, j := range records {
		// parse from file
		page := getPageFromFile(j.Url)
		////parse from url
		//page := getPage(j.Url)
		data := findData(page, j.Tag, j.Substring)
		if data != nil {
			fmt.Println(slices.Min(data))
		} else {
			fmt.Println("Empty")
		}

		fmt.Print("Press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
