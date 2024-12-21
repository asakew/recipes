package database

import (
	"log"

	"go-postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Distance struct {
	Db *gorm.DB
}

var DB Distance

// connectDb
func ConnectDb() {
	dsn := "host=37.77.104.186 user=gen_user password=o)\\;\\6!b,Tq^)= dbname=default_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&models.Book{})

	DB = Distance{
		Db: db,
	}
}
