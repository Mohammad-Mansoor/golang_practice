package handlers

import (
	"database/sql"
	"encoding/json"
	"time"

	// "log"
	"log/slog"
	"net/http"

	"github.com/Mohammad-Mansoor/go-api/internal/middlewares"
)

type listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
}
type ListingHandlers struct{
	db *sql.DB
}


func NewListingHandlers(db *sql.DB) *ListingHandlers{
	return &ListingHandlers{
		db: db,
	}
}

func (ls *ListingHandlers) Listing(w http.ResponseWriter, r *http.Request) {
	// this will prevent the zombie query 
		ctx := r.Context()
		requestID := middlewares.GetRequestId(ctx)
		slog.Info("listing.request", "request_id", requestID)

		rows, err := ls.db.QueryContext(ctx, `
		SELECT id, title, description, price, city, created_at FROM listings ORDER BY created_at DESC LIMIT 10
		`)

		if err != nil {
			slog.Error("listing.query", "error", err)
			http.Error(w, "Error occured during data fetching from Database ", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		listings := []listing{}
		for rows.Next() {

			var l listing
			if err := rows.Scan(&l.ID, &l.Title, &l.Description, &l.Price, &l.City, &l.CreatedAt); err != nil {
				slog.Error("rows.scan", "error", err)
				http.Error(w, "error while scanning the rows", http.StatusInternalServerError)
				return
			}
			listings = append(listings, l)
		}
		if err := rows.Err(); err != nil {
			slog.Error("rows.err", "error", err)
			http.Error(w, "row.error internal server: ", http.StatusInternalServerError)
			return
		}
		// log.Println(db)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(listings)

	}


func (ls *ListingHandlers) DeleteListing(w http.ResponseWriter, r *http.Request){

		id := r.PathValue("id")
		result, err := ls.db.Exec(`DELETE FROM listings WHERE id= $1`, id)
		if err!=nil{
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		affected, err := result.RowsAffected()
		if err != nil {
			slog.Error("result.rows_affected", "error", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		if affected == 0 {
			slog.Error("delete_listing", "error", "listing not found")
			http.Error(w, "listing not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
}