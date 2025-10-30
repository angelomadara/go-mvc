package routes

import (
	"mvc-app/controllers"

	"github.com/gorilla/mux"
)

// SetupRoutes configures all application routes
func SetupRoutes(userController *controllers.UserController) *mux.Router {
	router := mux.NewRouter()

	// User routes
	router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userController.GetUser).Methods("GET")
	router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	return router
}