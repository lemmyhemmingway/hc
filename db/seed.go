package db

import (
	"encoding/json"
	"log"
	"os"

	"healthcheck/models"
)

func SeedURLsFromFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Seed file not found (%s), skipping seeding: %v", path, err)
		return
	}
	defer file.Close()

	var urls []string
	err = json.NewDecoder(file).Decode(&urls)
	if err != nil {
		log.Fatalf("Failed to decode URL seed file: %v", err)
	}

	for _, u := range urls {
		var url models.URL
		result := DB.FirstOrCreate(&url, models.URL{Target: u})
		if result.Error != nil {
			log.Printf("Failed to insert URL %s: %v", u, result.Error)
		} else if result.RowsAffected > 0 {
			log.Printf("Seeded URL: %s", u)
		}
	}
}
