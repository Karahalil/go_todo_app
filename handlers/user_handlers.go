//Handler funnctions for the endpoints

package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/karahalil/backend-project/common"
	"github.com/karahalil/backend-project/db"
	"github.com/karahalil/backend-project/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Logic to get users from the database
	var users []models.User

	rows, err := db.DB.Query("SELECT id, username, email FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			common.WriteError(w, "Failed to scan user", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Logic to create a new user in the database
	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		common.WriteError(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if newUser.Username == "" || newUser.Email == "" {
		common.WriteError(w, "Username and email are required", http.StatusBadRequest)
		return
	}
	_, err := db.DB.Exec("INSERT INTO users (username, email) VALUES (?, ?)", newUser.Username, newUser.Email)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			common.WriteError(w, "Email already exists", http.StatusConflict)
			return
		}
		common.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Logic to delete a user from the database
	//Get the user ID from the url
	vars := mux.Vars(r)
	userID := vars["id"]
	_, err := db.DB.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		common.WriteError(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Logic to update a user's information in the database
	vars := mux.Vars(r)
	userID := vars["id"]
	var updatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		common.WriteError(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if updatedUser.Username == "" || updatedUser.Email == "" {
		common.WriteError(w, "Username and email are required", http.StatusBadRequest)
		return
	}
	_, err := db.DB.Exec("UPDATE users SET username = ?, email = ? WHERE id = ?", updatedUser.Username, updatedUser.Email, userID)
	if err != nil {
		common.WriteError(w, "Failed to update user", http.StatusInternalServerError)
		return
	}
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Logic to get a user by ID from the database
	vars := mux.Vars(r)
	userID := vars["id"]
	var user models.User
	row := db.DB.QueryRow("SELECT id, username, email FROM users WHERE id = ?", userID)
	if err := row.Scan(&user.ID, &user.Username, &user.Email); err != nil {
		if err.Error() == "sql: no rows in result set" {
			common.WriteError(w, "User not found", http.StatusNotFound)
		} else {
			common.WriteError(w, "Failed to retrieve user", http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(user)

}
