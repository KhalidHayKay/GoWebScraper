package main

import (
	"fmt"
	"log"
	"time"
	"web-scraper/config"
	"web-scraper/scraper"
)

func handleError(message string, err error) {
	log.Fatal(message + ": " + err.Error())
}

func main() {
	timeNow := time.Now()

	conf, err := config.Load()
	if err != nil {
		handleError("", err)
	}

	books := scraper.Run(conf)

	exporter := scraper.NewExporter(conf.PublicDir)
	if err := exporter.ToCSV(books); err != nil {
		handleError("Unable to export books", err)
	}
	if err := exporter.ToJSON(books); err != nil {
		handleError("Unable to export books", err)
	}

	fmt.Printf("Program ended after %v", time.Since(timeNow).String())
}
