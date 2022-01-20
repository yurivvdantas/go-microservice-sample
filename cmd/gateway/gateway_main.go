package main

import (
	"fmt"
	"go-microservice-sample/internal/gateway/infrastructure"
)

func main() {
	fmt.Println("Starting gateway server...")
	infrastructure.Dispatch()
}
