package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/niloysikdar/Sell-It/server/controllers"
	"github.com/niloysikdar/Sell-It/server/dbconnector"

	"github.com/gorilla/mux"
)

func main() {
	dbconnector.InitialMigration()
	initRouter()
	fmt.Println("Hello World")

}

func initRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/product/{id}", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/products", controllers.AddProduct).Methods("POST")
	router.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))

}
