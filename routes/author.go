package routes

import (
	"MobileApp/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

func RegisterHendler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Вставка пользователя в бд
		err = db.QueryRow(
			"INSERT INTO users (username, full_name, email) VALUES($1,$2,$3) RETURNING id",
			user.Username, user.Fullname, user.Email,
		).Scan(&user.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
