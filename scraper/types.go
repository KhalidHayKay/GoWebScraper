package scraper

import "strings"

type Book struct {
	Title   string
	Link    string
	InStock bool
	Price   float64
}

func (b *Book) getSlug() string {
	return strings.Split(b.Link, "/")[0]
}
