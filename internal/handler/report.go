package handler

import (
	"db-project/internal/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUserOrderSummary(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`
		SELECT u.ID, u.NAME, COALESCE(SUM(oi.QUANTITY * oi.UNIT_PRICE), 0) AS TOTAL_SPENT
		FROM USERS u
		LEFT JOIN ORDERS o ON u.ID = o.USER_ID
		LEFT JOIN ORDER_ITEMS oi ON o.ID = oi.ORDER_ID
		GROUP BY u.ID, u.NAME
		ORDER BY TOTAL_SPENT DESC
	`)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type UserSummary struct {
		ID         int     `json:"id"`
		Name       string  `json:"name"`
		TotalSpent float64 `json:"total_spent"`
	}

	var summaries []UserSummary
	for rows.Next() {
		var summary UserSummary
		if err := rows.Scan(&summary.ID, &summary.Name, &summary.TotalSpent); err != nil {
			http.Error(w, "Error parsing result", http.StatusInternalServerError)
			return
		}
		summaries = append(summaries, summary)
	}

	fmt.Println(summaries)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(summaries)
}

func GetBestSellingProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`
		SELECT 
			p.ID,
			p.NAME,
			SUM(oi.QUANTITY) AS TOTAL_SOLD,
			RANK() OVER (ORDER BY SUM(oi.QUANTITY) DESC) AS RANKING
		FROM PRODUCTS p
		INNER JOIN ORDER_ITEMS oi ON p.ID = oi.PRODUCT_ID
		GROUP BY p.ID, p.NAME
		ORDER BY TOTAL_SOLD DESC
	`)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type ProductRanking struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		TotalSold int    `json:"total_sold"`
		Ranking   int    `json:"ranking"`
	}

	var products []ProductRanking
	for rows.Next() {
		var pr ProductRanking
		if err := rows.Scan(&pr.ID, &pr.Name, &pr.TotalSold, &pr.Ranking); err != nil {
			http.Error(w, "Error parsing result", http.StatusInternalServerError)
			return
		}
		products = append(products, pr)
	}

	fmt.Println(products)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(products)
}

func GetProductsWithDetails(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`
		SELECT ID, NAME, DESCRIPTION, PRICE, DETAILS
		FROM PRODUCTS
	`)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Product struct {
		ID          int             `json:"id"`
		Name        string          `json:"name"`
		Description string          `json:"description"`
		Price       float64         `json:"price"`
		Details     json.RawMessage `json:"details"`
	}

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Details); err != nil {
			http.Error(w, "Error parsing result", http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	fmt.Println(products)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(products)
}
