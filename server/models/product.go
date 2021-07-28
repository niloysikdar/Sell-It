package models

import "gorm.io/gorm"

// Product model
type Product struct {
	gorm.Model
	ProductName string `json:"productName"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ImageURL    string `json:"imageurl"`
	Address     string `json:"address"`
	SellerName  string `json:"sellerName"`
	ContactInfo string `json:"contactInfo"`
	IsSold      bool   `json:"isSold"`
}
