package dbconnector

import (
	"fmt"

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
		fmt.Println(err.Error())
		panic("Failed to connect to DB !")
	}

	DB = db
	DB.AutoMigrate(&models.Product{})
	fmt.Println("DB Has been connected successfully ....")
}
