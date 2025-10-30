package services

import (
	"database/sql"
	"errors"
	"time"

	"mvc-app/models"
)

// UserService handles all user-related business logic and database operations
type UserService struct {
	db *sql.DB
}

// NewUserService creates a new user service
func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

// GetAllUsers retrieves all users from the database
func (s *UserService) GetAllUsers() ([]models.User, error) {
	rows, err := s.db.Query("SELECT id, name, email, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUserByID retrieves a user by ID with validation
func (s *UserService) GetUserByID(id int) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}

	var user models.User
	err := s.db.QueryRow("SELECT id, name, email, created_at, updated_at FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail retrieves a user by email
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.db.QueryRow("SELECT id, name, email, created_at, updated_at FROM users WHERE email = ?", email).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user with validation
func (s *UserService) CreateUser(user *models.User) error {
	// Validation
	if user.Name == "" {
		return errors.New("user name is required")
	}

	if user.Email == "" {
		return errors.New("user email is required")
	}

	// Check if email already exists
	existingUser, err := s.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return errors.New("user with this email already exists")
	}

	// Insert into database
	now := time.Now()
	result, err := s.db.Exec("INSERT INTO users (name, email, created_at, updated_at) VALUES (?, ?, ?, ?)",
		user.Name, user.Email, now, now)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	user.CreatedAt = now
	user.UpdatedAt = now
	return nil
}

// UpdateUser updates an existing user with validation
func (s *UserService) UpdateUser(user *models.User) error {
	if user.ID <= 0 {
		return errors.New("invalid user ID")
	}

	if user.Name == "" {
		return errors.New("user name is required")
	}

	if user.Email == "" {
		return errors.New("user email is required")
	}

	// Check if user exists
	existingUser, err := s.GetUserByID(user.ID)
	if err != nil {
		return err
	}

	if existingUser == nil {
		return errors.New("user not found")
	}

	// Check if email is taken by another user
	emailUser, err := s.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}

	if emailUser != nil && emailUser.ID != user.ID {
		return errors.New("email is already taken by another user")
	}

	// Update in database
	user.UpdatedAt = time.Now()
	_, err = s.db.Exec("UPDATE users SET name = ?, email = ?, updated_at = ? WHERE id = ?",
		user.Name, user.Email, user.UpdatedAt, user.ID)
	return err
}

// DeleteUser deletes a user by ID with validation
func (s *UserService) DeleteUser(id int) error {
	if id <= 0 {
		return errors.New("invalid user ID")
	}

	// Check if user exists
	existingUser, err := s.GetUserByID(id)
	if err != nil {
		return err
	}

	if existingUser == nil {
		return errors.New("user not found")
	}

	// Delete from database
	_, err = s.db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}