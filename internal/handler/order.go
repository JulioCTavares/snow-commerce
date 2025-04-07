package handler

import (
	"db-project/internal/db"
	"db-project/internal/model"
	"encoding/json"
	"net/http"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order model.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("INSERT INTO ORDERS (USER_ID, ORDER_DATE) VALUES (?, ?)", order.UserID, order.OrderDate)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT ID, USER_ID, ORDER_DATE FROM ORDERS")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var order model.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.OrderDate); err != nil {
			http.Error(w, "Error reading orders", http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
