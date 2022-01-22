package repository

import (
	"database/sql"
	"fmt"
	"go-microservice-sample/internal/crypto-votes-service/model"
)

func FindAllCrypto() ([]model.Cryptos, error) {
	conn, _ := InitConnection()
	var cryptos []model.Cryptos

	rows, err := conn.Query("SELECT * FROM CRYPTOS")
	if err != nil {
		return nil, fmt.Errorf("cryptoFindAll : %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var cry model.Cryptos
		if err := rows.Scan(&cry.Id, &cry.Name, &cry.Code, &cry.Upvote, &cry.Downvote, &cry.Description); err != nil {
			return nil, fmt.Errorf("cryptoFindAll : %v", err)
		}
		cryptos = append(cryptos, cry)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("cryptoFindAll : %v", err)
	}

	return cryptos, nil
}

func FindCryptoById(id int64) (*model.Cryptos, error) {
	conn, _ := InitConnection()
	var cryp model.Cryptos

	row := conn.QueryRow("SELECT * FROM CRYPTOS WHERE id = ?", id)

	if err := row.Scan(&cryp.Id, &cryp.Name, &cryp.Code, &cryp.Upvote, &cryp.Downvote, &cryp.Description); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("crypto %d: not found", id)
		}
		return nil, fmt.Errorf("cryptoById %d: %v", id, err)
	}

	return &cryp, nil
}

func AddCrypto(cry model.Cryptos) (int64, error) {
	conn, _ := InitConnection()
	result, err := conn.Exec("INSERT INTO CRYPTOS (name, code,upvotes,downvotes,description) VALUES (?, ?, ?,?,?)", cry.Name, cry.Code, 0, 0, cry.Description)
	if err != nil {
		return 0, fmt.Errorf("addCrypto: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addCrypto: %v", err)
	}
	return id, nil
}
