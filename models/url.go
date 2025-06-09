package models

type URL struct {
	ID     uint   `gorm:"primaryKey"`
	Target string `gorm:"uniqueIndex"`
}

