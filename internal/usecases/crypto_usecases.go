package usecases

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	pb "go-microservice-sample/api"
	"go-microservice-sample/configs"
	"go-microservice-sample/internal/model"
	"go-microservice-sample/internal/repository"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var db *sql.DB

func GetCryptoByID(ctxRequest *gin.Context) {
	id, err := strconv.ParseInt(ctxRequest.Param("id"), 10, 64)
	if err != nil {
		log.Fatalf("The id must be a int: %v", err)
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(configs.Microservice1Adress,
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

	cryptoById(id)

}

// cryptoById queries for albums that have the specified artist name.
func cryptoById(id int64) ([]model.Cryptos, error) {
	// An albums slice to hold data from returned rows.
	var cryptos []model.Cryptos

	rows, err := db.Query("SELECT * FROM CRYPTOS WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("cryptoById %q: %v", id, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var cry model.Cryptos
		if err := rows.Scan(&cry.Id, &cry.Name, &cry.Code, &cry.Upvote, &cry.Downvote, &cry.Description); err != nil {
			return nil, fmt.Errorf("cryptoById %q: %v", id, err)
		}
		cryptos = append(cryptos, cry)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", id, err)
	}
	return cryptos, nil
}
