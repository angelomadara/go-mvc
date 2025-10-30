package views

import (
	"encoding/json"
	"net/http"

	"mvc-app/models"
)

// UserView handles rendering user data
type UserView struct{}

// NewUserView creates a new user view
func NewUserView() *UserView {
	return &UserView{}
}

// RenderUsers renders a list of users as JSON
func (uv *UserView) RenderUsers(w http.ResponseWriter, users []models.User) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// RenderUser renders a single user as JSON
func (uv *UserView) RenderUser(w http.ResponseWriter, user models.User) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}