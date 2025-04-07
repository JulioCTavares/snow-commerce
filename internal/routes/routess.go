package routes

import (
	"db-project/internal/handler"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Users
	r.HandleFunc("/users", handler.CreateUser).Methods("POST")
	r.HandleFunc("/users", handler.GetUsers).Methods("GET")

	// Products
	r.HandleFunc("/products", handler.CreateProduct).Methods("POST")
	r.HandleFunc("/products", handler.GetProducts).Methods("GET")

	// Orders
	r.HandleFunc("/orders", handler.CreateOrder).Methods("POST")
	r.HandleFunc("/orders", handler.GetOrders).Methods("GET")

	// Order Items
	r.HandleFunc("/order-items", handler.CreateOrderItem).Methods("POST")
	r.HandleFunc("/order-items", handler.GetOrderItems).Methods("GET")

	// Relatórios
	r.HandleFunc("/users/order-summary", handler.GetUserOrderSummary).Methods("GET")
	r.HandleFunc("/products/top-selling", handler.GetBestSellingProducts).Methods("GET")
	r.HandleFunc("/orders/monthly-summary", handler.GetProductsWithDetails).Methods("GET")

	return r
}
