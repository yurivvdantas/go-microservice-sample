package usecases

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	pb "go-microservice-sample/api"
	"go-microservice-sample/configs"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func GetCryptoByID(ctxRequest *gin.Context) {
	id, err := strconv.ParseInt(ctxRequest.Param("id"), 10, 64)
	if err != nil {
		log.Fatalf("The id must be a int: %v", err)
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(configs.Crypto_votes_service_adress,
		grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUsersClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Get(ctx, &pb.UserId{Id: id})
	if err != nil {
		log.Fatalf("something get wrong on get user: %v", err)
	}
	log.Printf("Result: %s", r.Name)

	ctxRequest.IndentedJSON(http.StatusOK, r)
}
