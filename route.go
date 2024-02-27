package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", addProducts).Methods("POST")

	router.HandleFunc("/products", getAllProducts).Methods("GET")

	router.HandleFunc("/products/{id}", getProduct).Methods("GET")

	router.HandleFunc("/update/products/{id}", updateProduct).Methods("PUT")

	router.HandleFunc("/delete/products/{id}", deleteProduct).Methods("DELETE")

	router.HandleFunc("/category", addCategories).Methods("POST")

	router.HandleFunc("/categories", getAllCategories).Methods("GET")

	router.HandleFunc("/categories/{id}", getCategory).Methods("GET")

	router.HandleFunc("/update/categories/{id}", updateCategory).Methods("PUT")

	router.HandleFunc("/delete/categories/{id}", deleteCategory).Methods("DELETE")

	http.ListenAndServe(":5000", router)

}
