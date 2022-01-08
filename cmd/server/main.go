package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	pb "go-microservice-sample/api"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	router := gin.Default()
	router.GET("/users/:id", getUserByID)
	router.Run("localhost:8080")
}

func getUserByID(ctxRequest *gin.Context) {
	id, err := strconv.ParseInt(ctxRequest.Param("id"), 10, 64)
	if err != nil {
		log.Fatalf("The id must be a int: %v", err)
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
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
