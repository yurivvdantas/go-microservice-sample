package infrastructure

import (
	"go-microservice-sample/internal/gateway/usecases"

	"github.com/gin-gonic/gin"
)

func Dispatch() {
	router := gin.Default()
	router.GET("/crypto/:id", usecases.GetCryptoByID)
	router.GET("/crypto/", usecases.GetAllCrypto)
	router.POST("/crypto/", usecases.AddCrypto)
	router.POST("/crypto/:id/upvote", usecases.UpvoteCrypto)
	router.GET("/crypto/:id/upvote", usecases.GetUpvotes)
	router.Run("localhost:8080")
}
