package main

import (
	"go-microservice-sample/internal/infrastructure"
	"go-microservice-sample/internal/repository"

	"database/sql"
)

var db *sql.DB

const (
	address = "localhost:50051"
)

func main() {
	repository.InitConnection()
	infrastructure.Dispatch()
}
