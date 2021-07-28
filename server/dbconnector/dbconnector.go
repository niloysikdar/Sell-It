package dbconnector

import (
	"fmt"
	"log"

	"github.com/niloysikdar/Sell-It/server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB ...
var DB *gorm.DB

// InitialMigration ...
func InitialMigration() {
	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to DB !")
		log.Fatal(err.Error())
	}

	DB = db
	DB.AutoMigrate(&models.Product{})
	fmt.Println("DB Has been connected successfully ....")
}
