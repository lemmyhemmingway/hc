package models

import "time"

type HealthCheckRecord struct {
	ID           uint      `gorm:"primaryKey"`
	URLID        uint
	URL          URL       `gorm:"foreignKey:URLID"`
	StatusCode   int
	ResponseTime int64 // milliseconds
	Timestamp    time.Time
	Headers      string
}

