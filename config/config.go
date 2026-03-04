package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ScrapeURL string
	PublicDir string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		ScrapeURL: os.Getenv("SCRAPE_URL"),
		PublicDir: os.Getenv("PUBLIC_DIR"),
	}, nil
}
