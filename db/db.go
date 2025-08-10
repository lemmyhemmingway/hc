package db

import (
	"log"
	"os"

	"healthcheck/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=healthcheck port=5432 sslmode=disable"
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	err = DB.AutoMigrate(
		&models.Environment{},
		&models.Tag{},
		&models.URLTag{},
		&models.Location{},
		&models.CheckType{},
		&models.URL{},
		&models.HealthCheckRecord{},
		&models.UptimeStat{},
		&models.Type{},
		&models.Status{},
		&models.Target{},
		&models.Log{},
		&models.Schedule{},
	)
	if err != nil {
		log.Fatal("failed to migrate models:", err)
	}

	SeedURLsFromFile("db/urls.json")
}
