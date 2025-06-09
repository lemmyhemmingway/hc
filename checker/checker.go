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

func Check(urlStr string) {
	start := time.Now()
	resp, err := client.Get(urlStr)
	duration := time.Since(start).Milliseconds()

	if err != nil {
		log.Printf("[ERROR] %s - %s", urlStr, err)
		var url models.URL
		result := db.DB.FirstOrCreate(&url, models.URL{Target: urlStr})
		if result.Error == nil {
			record := models.HealthCheckRecord{
				URLID:        url.ID,
				StatusCode:   0,
				ResponseTime: duration,
				Timestamp:    time.Now(),
				Headers:      "",
			}
			db.DB.Create(&record)
		} else {
			log.Printf("[DB ERROR] could not create or fetch URL record: %v", result.Error)
		}
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

	result = db.DB.Create(&record)
	if result.Error != nil {
		log.Printf("[DB ERROR] could not create health check record: %v", result.Error)
		return
	}

	log.Printf("[INFO] %s - %d - %dms", urlStr, resp.StatusCode, duration)
}
