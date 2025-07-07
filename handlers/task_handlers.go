package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/karahalil/backend-project/common"
	"github.com/karahalil/backend-project/db"
	"github.com/karahalil/backend-project/models"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	vars := mux.Vars(r)
	userIDStr, ok := vars["id"]
	if !ok {
		common.WriteError(w, "user_id not found in URL", http.StatusBadRequest)
		return
	}
	user_id, err := strconv.Atoi(userIDStr)
	if err != nil {
		common.WriteError(w, "Invalid user_id format", http.StatusBadRequest)
		return
	}
	rows, err := db.DB.Query("SELECT id, title, description, status FROM tasks WHERE user_id = ?", user_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status); err != nil {
			common.WriteError(w, "Failed to scan task", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		common.WriteError(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	userIDStr, ok := body["user_id"]
	if !ok {
		common.WriteError(w, "user_id not found in request body", http.StatusBadRequest)
		return
	}
	user_id, err := strconv.Atoi(userIDStr)
	if err != nil {
		common.WriteError(w, "Invalid user_id format", http.StatusBadRequest)
		return
	}
	title := body["title"]
	description := body["description"]
	status := body["status"]
	if status == "" {
		status = "pending" // Default status if not provided
	}
	if title == "" || description == "" {
		common.WriteError(w, "Title, description are required", http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec("INSERT INTO tasks (title, description, status, user_id) VALUES (?, ?, ?, ?)", title, description, status, user_id)
	if err != nil {
		common.WriteError(w, "Failed to create task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task created successfully"})

}
