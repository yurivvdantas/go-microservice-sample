package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func InitConnection() (*sql.DB, error) {

	cfg := mysql.Config{
		User:                 "admin", //os.Getenv("DBUSER"),
		Passwd:               "admin", //os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "172.17.0.2:3306",
		DBName:               "cryptos",
		AllowNativePasswords: true,
	}

	// Get a database handle.
	var err error
	conn, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	pingErr := conn.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
		return nil, pingErr
	}
	fmt.Println("Connected!")

	return conn, nil
}
