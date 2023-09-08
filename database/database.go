package database

import (
	"go-middleware-recipe/model/core"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DBConn *gorm.DB
)

func ConnectDatabase() {

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected to database successfully.")

	_ = db.AutoMigrate(&core.User{})
	DBConn = db
}
