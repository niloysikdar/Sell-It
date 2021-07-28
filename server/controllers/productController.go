package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niloysikdar/Sell-It/server/dbconnector"
	"github.com/niloysikdar/Sell-It/server/models"
)

// GetProducts for getting all the products from DB
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []models.Product
	dbconnector.DB.Find(&products)
	json.NewEncoder(w).Encode(products)
}

// GetProduct for getting a particular product from DB
func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	var product models.Product
	dbconnector.DB.First(&product, id)
	json.NewEncoder(w).Encode(product)
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
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	var product models.Product
	dbconnector.DB.First(&product, id)
	var newProduct models.Product
	json.NewDecoder(r.Body).Decode(&newProduct)
	dbconnector.DB.Model(&product).Updates(newProduct)
	var updatedProduct models.Product
	dbconnector.DB.First(&updatedProduct, id)
	json.NewEncoder(w).Encode(updatedProduct)
}

// DeleteProduct for deleting a product from DB
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	var product models.Product
	dbconnector.DB.Delete(&product, id)
	w.Write([]byte("OK"))
}
