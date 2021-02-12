package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDatabaseClient() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=capstone-archive port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.Debug()

	err = db.AutoMigrate(&User{}, &Capstone{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
