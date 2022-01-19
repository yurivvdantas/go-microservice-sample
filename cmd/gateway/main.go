package main

import (
	"go-microservice-sample/internal/infrastructure"
	"go-microservice-sample/internal/repository"
)

func main() {
	repository.InitConnection()
	infrastructure.Dispatch()
}
