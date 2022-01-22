package repository

import (
	"database/sql"
	"fmt"
	"go-microservice-sample/configs"
	"log"

	"github.com/go-sql-driver/mysql"
)

func InitConnection() (*sql.DB, error) {

	cfg := mysql.Config{
		User:                 configs.User_database,     //os.Getenv("DBUSER"),
		Passwd:               configs.Password_database, //os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 configs.Address_database,
		DBName:               configs.Dbname_database,
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
	fmt.Println("Connected with database")

	return conn, nil
}
