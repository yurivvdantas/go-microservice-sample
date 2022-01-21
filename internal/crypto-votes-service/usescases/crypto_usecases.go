package usecases

import (
	"context"
	"log"

	pb "go-microservice-sample/api"
	"go-microservice-sample/internal/crypto-votes-service/repository"
)

// server is used to implement CryptoServer.GreeterServer.
type CryptoServer struct {
	pb.UnimplementedCryptosServer
}

// Get crypto by given id
func (s *CryptoServer) Get(ctx context.Context, in *pb.CryptoId) (*pb.Crypto, error) {
	log.Printf("Received userId: %v", in.Id)

	cryptos, err := repository.FindCryptoById(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Crypto{Id: cryptos.Id,
		Name:        cryptos.Name,
		Code:        cryptos.Code,
		Upvote:      cryptos.Upvote,
		Downvote:    cryptos.Downvote,
		Description: cryptos.Description}, nil
}
