package repositories

import (
	"connectopia-api/src/models"
	"database/sql"
)

/*
UsersRepository represents a repository for managing users.
It provides methods for interacting with the user data in the database.
*/
type UsersRepository struct {
	db *sql.DB
}

// NewUsersRepository creates a new instance of the UsersRepository.
// It takes a *sql.DB as a parameter and returns a pointer to the repository.
func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

// Create creates a new user in the database.
// It takes a user model as a parameter and returns the ID of the created user (uint64) and an error (if any).
func (repository *UsersRepository) Create(user models.UserModel) (uint64, error) {
	statment, err := repository.db.Prepare(
		"INSERT INTO users (name, username, email, password) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, nil
	}
	defer statment.Close()

	result, err := statment.Exec(user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	isertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(isertedID), nil
}
