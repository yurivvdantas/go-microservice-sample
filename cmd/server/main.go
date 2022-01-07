package main

import (
	"context"
	"log"
	"time"

	pb "go-microservice-sample/api"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUsersClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Get(ctx, &pb.UserId{Id: 1})
	if err != nil {
		log.Fatalf("something get wrong on get user: %v", err)
	}
	log.Printf("Result: %s", r.Name)

}
