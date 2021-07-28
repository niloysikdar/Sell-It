package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niloysikdar/Sell-It/server/controllers"
)

// InitRouter ...
func InitRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/product/{id}", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/products", controllers.AddProduct).Methods("POST")
	router.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Server is listening on PORT: 5000")
}
