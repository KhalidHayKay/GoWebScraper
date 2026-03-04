package scraper

import "strings"

type Book struct {
	Title   string  `json:"title"`
	Link    string  `json:"link"`
	InStock bool    `json:"in_stock"`
	Price   float64 `price:"price"`
}

func (b Book) Slug() string {
	return strings.Split(b.Link, "/")[0]
}
