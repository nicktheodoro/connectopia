package repositories

import (
	"connectopia-api/src/models"
	"database/sql"
	"fmt"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (repo *UsersRepository) Create(user models.UserModel) (uint64, error) {
	statement, err := repo.db.Prepare(
		"INSERT INTO users (name, username, email, password) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, nil
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(insertedID), nil
}

func (repo *UsersRepository) Update(ID int64, user models.UserModel) error {
	statement, err := repo.db.Prepare(
		"UPDATE users SET name = ?, username = ?, email = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Username, user.Email, ID); err != nil {
		return err
	}

	return nil
}

func (repo *UsersRepository) Delete(ID int64) error {
	statement, err := repo.db.Prepare(
		"DELETE FROM users WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repo UsersRepository) FindByNameOrUsername(nameOrUsername string) ([]models.UserModel, error) {
	nameOrUsername = fmt.Sprintf("%%%s%%", nameOrUsername)

	results, err := repo.db.Query(
		`SELECT 
				id, name, username, email, created_at, updated_at 
		FROM 
			users 
		WHERE 
			name LIKE ? OR username LIKE ?`,
		nameOrUsername, nameOrUsername,
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

func (repo UsersRepository) FindByEmail(email string) (models.UserModel, error) {
	result, err := repo.db.Query(
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

func (repo UsersRepository) GetUserByID(ID uint64) (models.UserModel, error) {
	result, err := repo.db.Query(
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

func (repo UsersRepository) GetUserPassword(ID uint64) (string, error) {
	row, err := repo.db.Query("SELECT password FROM users WHERE id = ?", ID)
	if err != nil {
		return "", err
	}
	defer row.Close()

	var user models.UserModel

	if row.Next() {
		if err = row.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (repo UsersRepository) GetFollowers(ID uint64) ([]models.UserModel, error) {
	results, err := repo.db.Query(
		`SELECT 
				u.id, u.name, u.username, u.email, u.created_at, u.updated_at 
		FROM 
			users u
			INNER JOIN followers f ON u.id = f.follower_id
		WHERE
			f.user_id = ?`,
		ID,
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

func (repo UsersRepository) GetFollowing(ID uint64) ([]models.UserModel, error) {
	results, err := repo.db.Query(
		`SELECT 
				u.id, u.name, u.username, u.email, u.created_at, u.updated_at 
		FROM 
			users u
			INNER JOIN followers f ON u.id = f.user_id
		WHERE
			f.follower_id = ?`,
		ID,
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

func (repo UsersRepository) Follow(followedID uint64, followerID uint64) error {
	statement, err := repo.db.Prepare(
		"INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(followedID, followerID); err != nil {
		return err
	}

	return nil
}

func (repo UsersRepository) Unfollow(unfollowedID uint64, unfollowerID uint64) error {
	statement, err := repo.db.Prepare(
		"DELETE FROM followers WHERE user_id = ? AND follower_id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(unfollowedID, unfollowerID); err != nil {
		return err
	}

	return nil
}

func (repo UsersRepository) UpdatePassword(userID uint64, password string) error {
	statement, err := repo.db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(password, userID); err != nil {
		return err
	}

	return nil
}
