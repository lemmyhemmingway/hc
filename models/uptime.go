package models

import "time"

// UptimeStat stores calculated uptime metrics for a URL.
type UptimeStat struct {
	ID              uint `gorm:"primaryKey"`
	URLID           uint `gorm:"uniqueIndex"`
	URL             URL  `gorm:"foreignKey:URLID"`
	Up              bool
	UptimePercent   float64
	UpSince         time.Time
	AvgResponseTime int64
	LastUpdated     time.Time
}
