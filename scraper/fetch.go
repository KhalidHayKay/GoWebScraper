package scraper

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 10,
}

func buildUrl(baseUrl string) []string {
	var urls []string

	for i := 1; i <= 50; i++ {
		url := baseUrl + "/catalogue/page-" + fmt.Sprint(i) + ".html"
		urls = append(urls, url)
	}

	return urls
}

func fetchPage(url string) ([]byte, error) {
	res, err := client.Get(url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("server response not OK")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
