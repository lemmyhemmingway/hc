package checker

import (
	"log"
	"net/http"
	"strings"
	"time"

	"healthcheck/db"
	"healthcheck/models"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func Check(u models.URL) {
	start := time.Now()
	resp, err := client.Get(u.Target)
	duration := time.Since(start).Milliseconds()

	if err != nil {
		log.Printf("[ERROR] %s - %s", u.Target, err)
		record := models.HealthCheckRecord{
			URLID:        u.ID,
			StatusCode:   0,
			ResponseTime: duration,
			Timestamp:    time.Now(),
			Headers:      "",
		}
		db.DB.Create(&record)
		return
	}
	defer resp.Body.Close()

	// Header'ları stringle
	var headers []string
	for k, v := range resp.Header {
		headers = append(headers, k+": "+strings.Join(v, ","))
	}

	record := models.HealthCheckRecord{
		URLID:        u.ID,
		StatusCode:   resp.StatusCode,
		ResponseTime: duration,
		Timestamp:    time.Now(),
		Headers:      strings.Join(headers, "; "),
	}

	result := db.DB.Create(&record)
	if result.Error != nil {
		log.Printf("[DB ERROR] could not create health check record: %v", result.Error)
		return
	}

	log.Printf("[INFO] %s - %d - %dms", u.Target, resp.StatusCode, duration)
}
