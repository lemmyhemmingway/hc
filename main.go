package main

import (
	"fmt"
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
	app.Get("/uptimes", getUptimes)

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

func getUptimes(c *fiber.Ctx) error {
	var urls []models.URL
	result := db.DB.Preload("Environment").Preload("Tags").Preload("Location").Preload("CheckType").Find(&urls)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	type UptimeItem struct {
		Up           bool     `json:"up"`
		Environment  string   `json:"environment"`
		Tags         []string `json:"tags"`
		Location     string   `json:"location"`
		Type         string   `json:"type"`
		Uptime       string   `json:"uptime"`
		UpSince      string   `json:"upSince"`
		ResponseTime string   `json:"responseTime"`
	}

	var items []UptimeItem

	for _, u := range urls {
		var records []models.HealthCheckRecord
		r := db.DB.Where("url_id = ?", u.ID).Order("timestamp desc").Limit(100).Find(&records)
		if r.Error != nil || len(records) == 0 {
			continue
		}

		var successCount int
		var totalResp int64
		upSince := records[0].Timestamp

		for i, rec := range records {
			if rec.StatusCode >= 200 && rec.StatusCode < 300 {
				successCount++
				if i == 0 {
					upSince = rec.Timestamp
				}
			} else if i == 0 {
				upSince = time.Time{}
			} else if upSince.Equal(records[i-1].Timestamp) {
				upSince = time.Time{}
			}
			totalResp += rec.ResponseTime
		}

		uptimePercent := float64(successCount) / float64(len(records)) * 100
		last := records[0]
		up := last.StatusCode >= 200 && last.StatusCode < 300
		avgResp := totalResp / int64(len(records))

		tagNames := make([]string, len(u.Tags))
		for i, t := range u.Tags {
			tagNames[i] = t.Name
		}

		upSinceStr := ""
		if !upSince.IsZero() {
			upSinceStr = upSince.Format("2006-01-02 15:04")
		}

		items = append(items, UptimeItem{
			Up:           up,
			Environment:  u.Environment.Name,
			Tags:         tagNames,
			Location:     u.Location.Name,
			Type:         u.CheckType.Name,
			Uptime:       fmt.Sprintf("%.1f%%", uptimePercent),
			UpSince:      upSinceStr,
			ResponseTime: fmt.Sprintf("%dms", avgResp),
		})
	}

	return c.JSON(items)
}
