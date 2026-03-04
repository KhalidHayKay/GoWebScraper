package scraper

import (
	"flag"
	"fmt"
	"log"
	"sync"
)

func getBooks(urls []string) []Book {
	// numberOfWorkers := 10
	numberOfWorkers := flag.Int("workers", 10, "nomber of concurent workers")
	flag.Parse()

	var jobs = make(chan string, *numberOfWorkers)
	var results = make(chan []Book, *numberOfWorkers)
	var errors = make(chan error, *numberOfWorkers)

	var wg sync.WaitGroup
	wg.Add(*numberOfWorkers)

	for range *numberOfWorkers {
		go worker(jobs, results, &wg, errors)
	}

	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(errors)
		close(results)
	}()

	go func() {
		for err := range errors {
			log.Printf("Worker error: %v", err)
		}
	}()

	var allBooks []Book

	for pageBook := range results {
		allBooks = append(allBooks, pageBook...)
	}

	return allBooks
}

func worker(jobs <-chan string, results chan<- []Book, wg *sync.WaitGroup, errors chan<- error) {
	defer wg.Done()

	for url := range jobs {
		page, err := fetchPage(url)
		if err != nil {
			errors <- fmt.Errorf("Unable to fetch %s: %w", url, err)
			continue
		}

		books, err := parse(page)
		if err != nil {
			errors <- fmt.Errorf("Unable to parse result of fetch %s: %w", url, err)
			continue
		}

		results <- books
	}
}
