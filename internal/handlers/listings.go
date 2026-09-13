package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	// "log"
	"net/http"
)

type listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
}

func Listing(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		rows, err := db.Query(`
		SELECT id, title, description, price, city, created_at FROM listings ORDER BY created_at DESC LIMIT 10
		`)

		if err != nil {
			http.Error(w, "Error occured during data fetching from Database ", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		listings := []listing{}
		for rows.Next() {

			var l listing
			if err := rows.Scan(&l.ID, &l.Title, &l.Description, &l.Price, &l.City, &l.CreatedAt); err != nil {
				log.Printf("rows.scan: %v", err)
				http.Error(w, "error while scanning the rows", http.StatusInternalServerError)
				return
			}
			listings = append(listings, l)
		}
		if err := rows.Err(); err != nil {
			log.Printf("rows.err: %v", err)
			http.Error(w, "row.error internal server: ", http.StatusInternalServerError)
			return
		}
		// log.Println(db)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(listings)

	}

}
