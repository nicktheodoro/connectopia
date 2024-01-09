package repositories

import (
	"connectopia-api/src/models"
	"database/sql"
	"fmt"
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

// Update updates an user in the database.
// It takes an ID and a user model as a parameter and returns error (if any).
func (repository *UsersRepository) Update(ID int64, user models.UserModel) error {
	statment, err := repository.db.Prepare(
		"UPDATE users set name = ?, username = ?, email = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statment.Close()

	if _, err := statment.Exec(user.Name, user.Username, user.Email, ID); err != nil {
		return err
	}

	return nil
}

// Delete deletes an user in the database.
// It takes an ID as a parameter and returns error (if any).
func (repository *UsersRepository) Delete(ID int64) error {
	statment, err := repository.db.Prepare(
		"DELETE FROM users WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statment.Close()

	if _, err := statment.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repository UsersRepository) FindByNameOrUsername(nameOrNick string) ([]models.UserModel, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	results, err := repository.db.Query(
		`SELECT 
				id, name, username, email, created_at, updated_at 
		FROM 
			users 
		WHERE 
			name LIKE ? or username LIKE ?`,
		nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var users []models.UserModel

	for results.Next() {
		var user models.UserModel

		if err = results.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository UsersRepository) FindByEmail(email string) (models.UserModel, error) {
	result, err := repository.db.Query(
		`SELECT 
				id, password 
		FROM 
			users 
		WHERE 
			email = ?`,
		email,
	)
	if err != nil {
		return models.UserModel{}, err
	}
	defer result.Close()

	var user models.UserModel

	for result.Next() {

		if err = result.Scan(&user.ID, &user.Password); err != nil {
			return models.UserModel{}, err
		}
	}

	return user, nil
}

// BuscarPorID traz um usuário do banco de dados
func (repository UsersRepository) GetUserByID(ID uint64) (models.UserModel, error) {
	result, err := repository.db.Query(
		`SELECT 
				id, name, username, email, created_at, updated_at 
		FROM 
			users 
		WHERE 
			id = ?`,
		ID,
	)
	if err != nil {
		return models.UserModel{}, err
	}
	defer result.Close()

	var user models.UserModel

	if result.Next() {
		if err = result.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return models.UserModel{}, err
		}
	}

	return user, nil
}
