package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Record struct {
	Name      string
	Url       string
	Tag       string
	Substring string
}

func readJson(records *[]Record) {
	f, err := os.Open("pagelist.json")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	jbody, err := io.ReadAll(f)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jbody, records)

	if err != nil {
		log.Fatal(err)
	}
}
