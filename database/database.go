package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() {
	dbUrl := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)
	log.Println(dbUrl)
	db, err := gorm.Open(
		postgres.Open(
			dbUrl,
		),
		&gorm.Config{},
	)
	if(err != nil){
		panic("cannot establish database connection")
	}
	log.Println("database connected successfully - ",db)
}