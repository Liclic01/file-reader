package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type JsonRepository interface {
	SaveJson(db *sql.DB, id string, hash string, content interface{}) error
}

type SqlJsonRepository struct {
}

func (ps *SqlJsonRepository) SaveJson(db *sql.DB, id string, hash string, content interface{}) error {
	contentJSON, err := json.Marshal(content)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %v", err)
	}

	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM file_content WHERE hash = $1", hash)
	err = row.Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check existing record: %v", err)
	}

	if count == 0 {
		query := `
			INSERT INTO file_content (id, hash, content)
			VALUES ($1, $2, $3);
		`

		_, err = db.Exec(query, id, hash, contentJSON)
		if err != nil {
			return fmt.Errorf("failed to save data to PostgreSQL: %v", err)
		}
	}

	return nil
}
