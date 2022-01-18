package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	pb "go-microservice-sample/api"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"database/sql"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

const (
	address     = "localhost:50051"
	defaultName = "world"
)

type Cryptos struct {
	Id          int64
	Name        string
	Code        string
	Upvote      int64
	Downvote    int64
	Description string
}

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

	initConnectionWithDatabase()

}

func initConnectionWithDatabase() {
	cfg := mysql.Config{
		User:                 "admin", //os.Getenv("DBUSER"),
		Passwd:               "admin", //os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "172.17.0.2:3306",
		DBName:               "cryptos",
		AllowNativePasswords: true,
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	cryptos, err := cryptoById(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cryptos found: %v\n", cryptos)
}

// cryptoById queries for albums that have the specified artist name.
func cryptoById(id int64) ([]Cryptos, error) {
	// An albums slice to hold data from returned rows.
	var cryptos []Cryptos

	rows, err := db.Query("SELECT * FROM CRYPTOS WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("cryptoById %q: %v", id, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var cry Cryptos
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
