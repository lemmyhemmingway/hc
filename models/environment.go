package models

// Environment represents a deployment environment like prod or nonprod.
type Environment struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex"`
}
