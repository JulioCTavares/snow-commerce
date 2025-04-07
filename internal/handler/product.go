package handler

import (
	"db-project/internal/db"
	"db-project/internal/model"
	"encoding/json"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("INSERT INTO PRODUCTS (NAME, DESCRIPTION, PRICE) VALUES (?, ?, ?)", product.Name, product.Description, product.Price)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT ID, NAME, DESCRIPTION, PRICE FROM PRODUCTS")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price); err != nil {
			http.Error(w, "Error reading products", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
