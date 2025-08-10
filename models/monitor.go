package models

import "time"

// Type represents the type of a monitored target (e.g., domain or IP).
type Type struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex;size:50;not null"`
}

// Status represents a log status code.
type Status struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex;size:30;not null"`
}

// Target represents a monitored domain or IP address.
type Target struct {
	ID        uint   `gorm:"primaryKey"`
	Domain    string `gorm:"size:255"`
	IPAddress string `gorm:"type:inet"`
	TypeID    uint   `gorm:"not null"`
	Type      Type   `gorm:"foreignKey:TypeID"`
}

// Log represents a monitoring log entry.
type Log struct {
	ID        uint      `gorm:"primaryKey"`
	TargetID  uint      `gorm:"not null"`
	Target    Target    `gorm:"foreignKey:TargetID;constraint:OnDelete:CASCADE"`
	StatusID  uint      `gorm:"not null"`
	Status    Status    `gorm:"foreignKey:StatusID"`
	Timestamp time.Time `gorm:"autoCreateTime"`
	Message   string
}

// Schedule represents a monitoring schedule for a target.
type Schedule struct {
	ID              uint      `gorm:"primaryKey"`
	TargetID        uint      `gorm:"not null"`
	Target          Target    `gorm:"foreignKey:TargetID;constraint:OnDelete:CASCADE"`
	IntervalSeconds int       `gorm:"not null"`
	NextRunAt       time.Time `gorm:"not null"`
	Enabled         bool      `gorm:"default:true"`
}
