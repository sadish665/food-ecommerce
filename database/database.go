package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() {
	db, err := gorm.Open(
		postgres.Open(
			"host=localhost port=5433 user=sadish dbname=food_ecommerce_db password=1234 sslmode=disable",
		),
		&gorm.Config{},
	)
	if(err != nil){
		panic("cannot establish database connection")
	}
	log.Println("database connected successfully - ",db)
}