package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DBConn *gorm.DB
)

func ConnectDatabase() {
	dsn := "root:root@tcp(127.0.0.1:33060)/go_middleware_recipe?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected to database successfully.")

	// TODO: BA
	//	db.AutoMigrate(&models.Book{})
	DBConn = db
}
