package usecases

import (
	"context"
	"log"

	pb "go-microservice-sample/api"
	"go-microservice-sample/internal/crypto-votes-service/repository"

	"github.com/golang/protobuf/ptypes/empty"
)

// server is used to implement CryptoServer.GreeterServer.
type CryptoServer struct {
	pb.UnimplementedCryptosServer
}

// Get crypto by given id
func (s *CryptoServer) Get(ctx context.Context, in *pb.CryptoId) (*pb.Crypto, error) {
	log.Printf("Received cryptoId: %v", in.Id)

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

func (s *CryptoServer) GetAll(in *empty.Empty, srv pb.Cryptos_GetAllServer) error {
	log.Printf("Getting all cryptos...")

	cryptos, err := repository.FindAllCrypto()
	if err != nil {
		return err
	}
	for _, c := range cryptos {
		resp := pb.Crypto{Id: c.Id,
			Name:        c.Name,
			Code:        c.Code,
			Upvote:      c.Upvote,
			Downvote:    c.Downvote,
			Description: c.Description}
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
		log.Printf("sending data: %v", c.Code)
	}
	return nil
}
