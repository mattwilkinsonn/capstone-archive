package db

import (
	"log"
	"os"

	"github.com/Zireael13/capstone-archive/server/internal/envs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitalizeDatabase(orm *gorm.DB) {
	err := orm.AutoMigrate(&User{}, &Capstone{})

	if err != nil {
		log.Fatal(err)
	}

}

func CreateDatabaseDialector() gorm.Dialector {
	dsn := "host=localhost user=postgres password=postgres dbname=capstone-archive port=5432 sslmode=disable"
	return postgres.Open(dsn)
}

func CreateDatabaseClient(dialector gorm.Dialector) *gorm.DB {
	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	env := envs.GetEnvironment()
	if env == "development" {
		db.Debug()
	}

	return db
}

func CreateDefaultDatabaseClient() *gorm.DB {
	dia := CreateDatabaseDialector()
	orm := CreateDatabaseClient(dia)
	InitalizeDatabase(orm)
	return orm
}
func LoadSampleData(orm *gorm.DB) {

	var capstone Capstone

	res := orm.First(&capstone)

	if res.Error != nil {
		file, err := os.ReadFile("../sample/capstones.sql")

		if err != nil {
			panic(err)
		}

		orm.Exec(string(file))

	}
}
