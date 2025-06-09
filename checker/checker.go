package checker

import (
	"log"
	"net/http"
	"strings"
	"time"

	"healthcheck/db"
	"healthcheck/models"
)

func Check(urlStr string) {
	start := time.Now()
	resp, err := http.Get(urlStr)
	duration := time.Since(start).Milliseconds()

	if err != nil {
		log.Printf("[ERROR] %s - %s", urlStr, err)
		return
	}
	defer resp.Body.Close()

	// URL tablosundan ID’yi al veya ekle
	var url models.URL
	result := db.DB.FirstOrCreate(&url, models.URL{Target: urlStr})
	if result.Error != nil {
		log.Printf("[DB ERROR] could not create or fetch URL record: %v", result.Error)
		return
	}

	// Header'ları stringle
	var headers []string
	for k, v := range resp.Header {
		headers = append(headers, k+": "+strings.Join(v, ","))
	}

	record := models.HealthCheckRecord{
		URLID:        url.ID,
		StatusCode:   resp.StatusCode,
		ResponseTime: duration,
		Timestamp:    time.Now(),
		Headers:      strings.Join(headers, "; "),
	}

	db.DB.Create(&record)

	log.Printf("[INFO] %s - %d - %dms", urlStr, resp.StatusCode, duration)
}

