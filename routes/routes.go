//Url routes for the application

package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/karahalil/backend-project/handlers"
)

func Setup() *mux.Router {
	router := mux.NewRouter()
	// User routes
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
	//Task routes
	router.HandleFunc("/tasks/{id}", handlers.GetTasks).Methods("POST")
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	//Frontend routes
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))
	return router
}
