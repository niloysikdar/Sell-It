package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/niloysikdar/Sell-It/server/dbconnector"
	"github.com/niloysikdar/Sell-It/server/models"
)

// GetProducts for getting all the products from DB
func GetProducts(w http.ResponseWriter, r *http.Request) {

}

// GetProduct for getting a particular product from DB
func GetProduct(w http.ResponseWriter, r *http.Request) {

}

// AddProduct for adding a product to DB
func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	dbconnector.DB.Create(&product)
	json.NewEncoder(w).Encode(product)
}

// UpdateProduct for updating a particular product in DB
func UpdateProduct(w http.ResponseWriter, r *http.Request) {

}

// DeleteProduct for deleting a product from DB
func DeleteProduct(w http.ResponseWriter, r *http.Request) {

}
