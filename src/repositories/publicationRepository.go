package repositories

import (
	"connectopia-api/src/models"
	"database/sql"
)

// PublicationRepository represents a repository of publications
type PublicationRepository struct {
	db *sql.DB
}

// NewPublicationRepository creates a repository for publications
func NewPublicationRepository(db *sql.DB) *PublicationRepository {
	return &PublicationRepository{db}
}

// Get fetches publications from followed users and the requesting user
func (repository PublicationRepository) Get(userID uint64) ([]models.PublicationModel, error) {
	rows, err := repository.db.Query(`
		SELECT DISTINCT p.*, u.username as authorNick FROM publications p
		INNER JOIN users u ON u.id = p.author_id
		LEFT JOIN followers f ON p.author_id = f.user_id
		WHERE u.id = ? OR f.follower_id = ?
		ORDER BY 1 DESC`,
		userID, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.PublicationModel

	for rows.Next() {
		var publication models.PublicationModel

		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.UpdateAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// GetByID fetches a single publication from the database
func (repository PublicationRepository) GetByID(publicationID uint64) (models.PublicationModel, error) {
	row, err := repository.db.Query(`
		SELECT p.*, u.username as authorNick FROM publications p
		INNER JOIN users u ON u.id = p.author_id WHERE p.id = ?`,
		publicationID,
	)
	if err != nil {
		return models.PublicationModel{}, err
	}
	defer row.Close()

	var publication models.PublicationModel

	if row.Next() {
		if err = row.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.UpdateAt,
			&publication.AuthorNick,
		); err != nil {
			return models.PublicationModel{}, err
		}
	}

	return publication, nil
}

// GetByUser fetches all publications of a specific user
func (repository PublicationRepository) GetByUser(userID uint64) ([]models.PublicationModel, error) {
	rows, err := repository.db.Query(`
		SELECT p.*, u.username as authorNick FROM publications p
		JOIN users u ON u.id = p.author_id
		WHERE p.author_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.PublicationModel

	for rows.Next() {
		var publication models.PublicationModel

		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.UpdateAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// Create inserts a publication into the database
func (repository PublicationRepository) Create(publication models.PublicationModel) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO publications (title, content, author_id) VALUES (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(publication.Title, publication.Content, publication.AuthorID)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

// Update modifies the data of a publication in the database
func (repository PublicationRepository) Update(publicationID uint64, publication models.PublicationModel) error {
	statement, err := repository.db.Prepare("UPDATE publications SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publication.Title, publication.Content, publicationID); err != nil {
		return err
	}

	return nil
}

// Delete deletes a publication from the database
func (repository PublicationRepository) Delete(publicationID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM publications WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}

// Like adds a like to the publication
func (repository PublicationRepository) Like(publicationID uint64) error {
	statement, err := repository.db.Prepare("UPDATE publications SET likes = likes + 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}

// Unlike subtracts a like from the publication
func (repository PublicationRepository) Unlike(publicationID uint64) error {
	statement, err := repository.db.Prepare(`
		UPDATE publications SET likes = 
		CASE 
			WHEN likes > 0 THEN likes - 1
			ELSE 0 
		END
		WHERE id = ?
	`)
	if err != nil {
		return err
	}

	if _, err = statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}
