package config

import (
	"log"
	"os"

	"emad.com/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDB()  {
	db, err := gorm.Open("mysql",os.Getenv("DB_URL"))

	if err != nil {

		log.Fatalf("Could not connect to database: %s", err)
	}

	db.AutoMigrate(
		&models.Users{}, 
		&models.Auths{},
	)
	DB = db
	
}