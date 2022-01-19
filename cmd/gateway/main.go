package main

import (
	"go-microservice-sample/internal/infrastructure"

	"database/sql"
)

var db *sql.DB

const (
	address = "localhost:50051"
)

func main() {
	infrastructure.Dispatch()
}
