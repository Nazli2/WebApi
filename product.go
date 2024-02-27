package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var products []Product = []Product{}

type Product struct {
	ID          int     `json:"id"`
	ProductName string  `json:"productName"`
	UnitPrice   float64 `json:"unitPrice"`
	CategoryID  int     `json:"categoryId"`
}

func addProducts(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	products = append(products, newProduct)

	json.NewEncoder(w).Encode(products)
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to Integer"))
		return
	}

	if id >= len(products) {
		w.WriteHeader(404)
		w.Write([]byte("no profile found with specified ID"))
		return
	}
	product := products[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}

func updateProduct(w http.ResponseWriter, r *http.Request) {

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to Integer"))
		return
	}

	if id >= len(products) {
		w.WriteHeader(404)
		w.Write([]byte("no profile found with specified ID"))
		return
	}

	var updateProduct Product
	json.NewDecoder(r.Body).Decode(&updateProduct)

	products[id] = updateProduct

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateProduct)

}

func deleteProduct(w http.ResponseWriter, r *http.Request) {

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to Integer"))
		return
	}

	if id >= len(products) {
		w.WriteHeader(404)
		w.Write([]byte("no profile found with specified ID"))
		return
	}

	products = append(products[:id], products[:id+1]...)

	w.WriteHeader(200)

}
