package dbconnector

import (
	"fmt"

	"github.com/niloysikdar/Sell-It/server/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitialMigration ...
func InitialMigration() {
	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DB !")
	}

	db.AutoMigrate(&models.Product{})
	fmt.Println("DB Has been connected successfully")
}
