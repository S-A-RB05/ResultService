package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Open the HTM file
	file, err := os.Open("tester_report.htm")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Load the HTM document
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(err)
	}

	// Print all the extracted data from the HTM
	fmt.Println(doc.Text())
}
