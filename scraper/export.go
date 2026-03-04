package scraper

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type Exporter struct {
	OutputDir string
}

func NewExporter(outputDir string) *Exporter {
	return &Exporter{OutputDir: outputDir}
}

func (b *Exporter) ToCSV(books []Book) error {
	file, err := os.Create(b.OutputDir + "/books.csv")

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"s/n", "Name", "Slug", "In Stock", "Price (£)"})

	for i, book := range books {
		writer.Write([]string{
			fmt.Sprint(i + 1),
			book.Title,
			book.Slug(),
			fmt.Sprint(book.InStock),
			fmt.Sprint(book.Price),
		})
	}

	return nil
}

func (d *Exporter) ToJSON(books []Book) error {
	file, err := os.Create(d.OutputDir + "/books.json")
	if err != nil {
		return nil
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(books)

	return nil
}
