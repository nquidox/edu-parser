package main

import (
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func getPage(url string) *html.Node {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}
func getPageFromFile(path string) *html.Node {
	f, err := os.Open(path)
	if err != nil {
		log.Println("Error opening file", err)
	}
	defer f.Close()

	doc, err := html.Parse(f)

	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func findData(doc *html.Node, tag, substring string) []int {
	var (
		crawler func(*html.Node)
		values  []int
	)

	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == tag {
			if strings.Contains(node.FirstChild.Data, substring) {
				values = append(values, stripNum(node.FirstChild.Data))
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	return values
}

func stripNum(num string) int {
	number := ""
	for _, j := range num {
		if j >= '0' && j <= '9' {
			number += string(j)
		}
	}

	n, err := strconv.Atoi(number)

	if err != nil {
		log.Fatal(err)
	}

	return n
}
