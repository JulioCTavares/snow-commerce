package handler

import (
	"db-project/internal/db"
	"db-project/internal/model"
	"encoding/json"
	"net/http"
)

func CreateOrderItem(w http.ResponseWriter, r *http.Request) {
	var item model.OrderItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("INSERT INTO ORDER_ITEMS (PRODUCT_ID, QUANTITY, UNIT_PRICE) VALUES (?, ?, ?)", item.ProductID, item.Quantity, item.UnitPrice)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetOrderItems(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT ID, PRODUCT_ID, QUANTITY, UNIT_PRICE FROM ORDER_ITEMS")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []model.OrderItem
	for rows.Next() {
		var item model.OrderItem
		if err := rows.Scan(&item.ID, &item.ProductID, &item.Quantity, &item.UnitPrice); err != nil {
			http.Error(w, "Error reading order items", http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
