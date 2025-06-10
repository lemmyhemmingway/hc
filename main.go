package main

import (
	"log"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"

	"healthcheck/checker"
	"healthcheck/db"
	"healthcheck/models"
)

func main() {
	db.Init()

	go startHealthCheckLoop()

	app := fiber.New()

	app.Get("/urls", getURLs)
	app.Get("/records", getRecords)

	app.Static("/", "./ui/dist/ui/browser")
	app.Use(func(c *fiber.Ctx) error {
		return c.SendFile("./ui/dist/ui/browser/index.html")
	})

	log.Fatal(app.Listen(":3000"))
}

func startHealthCheckLoop() {
	interval := 30 * time.Second
	for {
		var urls []models.URL
		result := db.DB.Preload("Environment").Preload("Tags").Preload("Location").Preload("CheckType").Find(&urls)
		if result.Error != nil {
			log.Printf("Failed to load URLs from database: %v", result.Error)
			time.Sleep(interval)
			continue
		}

		var wg sync.WaitGroup
		wg.Add(len(urls))

		for _, url := range urls {
			urlCopy := url
			go func(u models.URL) {
				defer wg.Done()
				checker.Check(u)
			}(urlCopy)
		}

		wg.Wait()
		time.Sleep(interval)
	}
}

func getURLs(c *fiber.Ctx) error {
	var urls []models.URL
	result := db.DB.Preload("Environment").Preload("Tags").Preload("Location").Preload("CheckType").Find(&urls)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}
	return c.JSON(urls)
}

func getRecords(c *fiber.Ctx) error {
	var records []models.HealthCheckRecord
	result := db.DB.Preload("URL.Environment").Preload("URL.Tags").Preload("URL.Location").Preload("URL.CheckType").Order("timestamp desc").Find(&records)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}
	return c.JSON(records)
}
