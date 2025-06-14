package db

import (
	"log"

	"healthcheck/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("healthcheck.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	err = DB.AutoMigrate(&models.Environment{}, &models.Tag{}, &models.URLTag{}, &models.Location{}, &models.CheckType{}, &models.URL{}, &models.HealthCheckRecord{}, &models.UptimeStat{})
	if err != nil {
		log.Fatal("failed to migrate models:", err)
	}

	SeedURLsFromFile("db/urls.json")
}
