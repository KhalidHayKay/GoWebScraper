package scraper

import "web-scraper/config"

func Run(conf *config.Config) []Book {
	urls := buildUrl(conf.ScrapeURL)

	books := getBooks(urls)

	return books
}
