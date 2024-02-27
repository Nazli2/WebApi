package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var categories []Category = []Category{}

type Category struct {
	ID           int    `json: "id"`
	CategoryName string `json: "categoryName"`
}

func addCategories(w http.ResponseWriter, r *http.Request) {
	var newCategory Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	categories = append(categories, newCategory)

	json.NewEncoder(w).Encode(categories)

}

func getAllCategories(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(categories)
}

func getCategory(w http.ResponseWriter, r *http.Request) {

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to Integer"))
		return
	}

	if id >= len(categories) {
		w.WriteHeader(404)
		w.Write([]byte("no profile found with specified ID"))
		return
	}
	category := categories[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)

}

func updateCategory(w http.ResponseWriter, r *http.Request) {

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to Integer"))
		return
	}

	if id >= len(categories) {
		w.WriteHeader(404)
		w.Write([]byte("no profile found with specified ID"))
		return
	}

	var updateCategory Category
	json.NewDecoder(r.Body).Decode(&updateCategory)

	categories[id] = updateCategory

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateCategory)

}

func deleteCategory(w http.ResponseWriter, r *http.Request) {

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to Integer"))
		return
	}

	if id >= len(categories) {
		w.WriteHeader(404)
		w.Write([]byte("no profile found with specified ID"))
		return
	}

	categories = append(categories[:id], categories[:id+1]...)

	w.WriteHeader(200)

}
