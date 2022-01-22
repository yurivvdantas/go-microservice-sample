package infrastructure

import (
	"go-microservice-sample/internal/gateway/usecases"

	"github.com/gin-gonic/gin"
)

func Dispatch() {
	router := gin.Default()
	router.GET("/users/:id", usecases.GetCryptoByID)
	router.GET("/users/", usecases.GetAllCrypto)
	router.Run("localhost:8080")
}
