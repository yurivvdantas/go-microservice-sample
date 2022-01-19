package repository

import (
	"database/sql"
	"fmt"
	"go-microservice-sample/internal/model"
)

type CrypotRepository struct {
	db *sql.DB
}

func (cr *CrypotRepository) FindById(id int64) ([]model.Cryptos, error) {
	// An albums slice to hold data from returned rows.
	var cryptos []model.Cryptos

	rows, err := cr.db.Query("SELECT * FROM CRYPTOS WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("cryptoById %q: %v", id, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var cry model.Cryptos
		if err := rows.Scan(&cry.Id, &cry.Name, &cry.Code, &cry.Upvote, &cry.Downvote, &cry.Description); err != nil {
			return nil, fmt.Errorf("cryptoById %q: %v", id, err)
		}
		cryptos = append(cryptos, cry)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", id, err)
	}

	return cryptos, nil
}
