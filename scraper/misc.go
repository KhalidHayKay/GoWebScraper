package scraper

import (
	"strconv"
	"strings"
)

func priceToFloat(price string) (float64, error) {
	clean := strings.TrimSpace(price)
	clean = strings.ReplaceAll(clean, "£", "")
	clean = strings.ReplaceAll(clean, ",", "")

	return strconv.ParseFloat(clean, 64)
}
