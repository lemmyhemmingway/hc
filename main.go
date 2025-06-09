package main

import (
	"log"
	"sync"
	"time"

	"healthcheck/checker"
	"healthcheck/db"
	"healthcheck/models"
)

func main() {
	db.Init()

	interval := 30 * time.Second

	for {
		var urls []models.URL
		result := db.DB.Find(&urls)
		if result.Error != nil {
			log.Printf("Failed to load URLs from database: %v", result.Error)
			continue
		}

		var wg sync.WaitGroup
		wg.Add(len(urls))

		for _, url := range urls {
			go func(u string) {
				defer wg.Done()
				checker.Check(u)
			}(url.Target)
		}

		wg.Wait()
		time.Sleep(interval)
	}
}

