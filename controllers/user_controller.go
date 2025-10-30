package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"mvc-app/models"
	"mvc-app/services"
	"mvc-app/views"

	"github.com/gorilla/mux"
)

// UserController handles HTTP requests for users
type UserController struct {
	userService *services.UserService
	userView    *views.UserView
}

// NewUserController creates a new user controller
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
		userView:    views.NewUserView(),
	}
}

// GetUsers handles GET /users
func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	uc.userView.RenderUsers(w, users)
}

// GetUser handles GET /users/{id}
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := uc.userService.GetUserByID(id)
	if err != nil {
		if err.Error() == "user not found" {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	uc.userView.RenderUser(w, *user)
}

// CreateUser handles POST /users
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := uc.userService.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	uc.userView.RenderUser(w, user)
}

// UpdateUser handles PUT /users/{id}
func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user.ID = id
	if err := uc.userService.UpdateUser(&user); err != nil {
		if err.Error() == "user not found" {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	uc.userView.RenderUser(w, user)
}

// DeleteUser handles DELETE /users/{id}
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := uc.userService.DeleteUser(id); err != nil {
		if err.Error() == "user not found" {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}