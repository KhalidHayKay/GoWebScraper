package scraper

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func parse(html []byte) ([]Book, error) {
	document, err := goquery.NewDocumentFromReader(bytes.NewReader(html))

	if err != nil {
		return nil, err
	}

	var books []Book

	document.Find(".product_pod").Each(func(i int, s *goquery.Selection) {
		title, _ := s.Find(" h3 a").First().Attr("title")
		href, _ := s.Find(" h3 a").First().Attr("href")

		priceAbs := s.Find("div.product_price p.price_color").Text()
		availability := s.Find("div.product_price p.availability").Text()

		price, _ := priceToFloat(priceAbs)

		books = append(books, Book{
			Title:   title,
			Link:    href,
			InStock: strings.TrimSpace(availability) == "In stock",
			Price:   price,
		})
	})

	return books, nil
}
