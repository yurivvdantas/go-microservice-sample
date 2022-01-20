package usecases

import (
	"context"
	"fmt"
	"log"

	pb "go-microservice-sample/api"
	"go-microservice-sample/internal/crypto-votes-service/repository"
)

// server is used to implement helloworld.GreeterServer.
type CryptoServer struct {
	pb.UnimplementedUsersServer
}

// SayHello implements helloworld.GreeterServer
func (s *CryptoServer) Get(ctx context.Context, in *pb.UserId) (*pb.User, error) {
	log.Printf("Received userId: %v", in.Id)

	cryptos, _ := repository.FindCryptoById(in.Id)
	fmt.Println(cryptos)

	return &pb.User{Id: in.Id, Name: "New user"}, nil
}
