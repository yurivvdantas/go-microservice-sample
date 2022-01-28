package usecases

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "go-microservice-sample/api"
	"go-microservice-sample/internal/crypto-votes-service/cache"
	"go-microservice-sample/internal/crypto-votes-service/model"
	"go-microservice-sample/internal/crypto-votes-service/repository"

	"github.com/golang/protobuf/ptypes/empty"
)

// server is used to implement CryptoServer.GreeterServer.
type CryptoServer struct {
	pb.UnimplementedCryptosServer
}

// Get crypto by given id
func (s *CryptoServer) Get(ctx context.Context, in *pb.CryptoId) (*pb.Crypto, error) {

	log.Printf("Received cryptoId: %v", in.GetId())

	cryCache := cache.GetCache(fmt.Sprintf("%d", in.GetId()))

	if cryCache != nil {
		fmt.Println("Getting crytpo from cache")
		return &pb.Crypto{Id: cryCache.Id,
			Name:        cryCache.Name,
			Code:        cryCache.Code,
			Upvote:      cryCache.Upvote,
			Downvote:    cryCache.Downvote,
			Description: cryCache.Description}, nil
	}

	cryptos, err := repository.FindCryptoById(in.GetId())
	if err != nil {
		return nil, err
	}

	cache.SetCache(fmt.Sprintf("%d", cryptos.Id), cryptos)
	fmt.Println("Updating cache")

	return &pb.Crypto{Id: cryptos.Id,
		Name:        cryptos.Name,
		Code:        cryptos.Code,
		Upvote:      cryptos.Upvote,
		Downvote:    cryptos.Downvote,
		Description: cryptos.Description}, nil
}

// Get all crypto stored in database
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

// Add crypto by given name, description and code
func (s *CryptoServer) AddCrypto(ctx context.Context, in *pb.Crypto) (*pb.CryptoId, error) {
	log.Printf("Received crypto name: %v", in.Name)

	id, err := repository.AddCrypto(model.Cryptos{Name: in.Name, Description: in.Description, Code: in.Code})
	if err != nil {
		return nil, err
	}

	return &pb.CryptoId{Id: id}, nil
}

// Increse by one a upvote to a crypt
func (s *CryptoServer) Upvote(ctx context.Context, in *pb.CryptoId) (*empty.Empty, error) {
	log.Printf("Received crypto id: %v", in)

	cache.Clean(fmt.Sprintf("%d", in.GetId()))
	log.Printf("Cache clean")

	crypto, err := repository.FindCryptoById(in.GetId())
	if err != nil {
		return nil, err
	}
	crypto.Upvote += 1

	return &empty.Empty{}, repository.UpdateCrypto(crypto)
}

// Stream upvotes of a crypto
func (s *CryptoServer) LiveUpVotes(in *pb.CryptoId, srv pb.Cryptos_LiveUpVotesServer) error {
	log.Printf("Initialize a stream of upvotes...")

	for {
		time.Sleep(time.Second)
		upvotes, err := repository.FindUpVotesCryptoById(in.GetId())
		if err != nil {
			return err
		}

		if err := srv.Send(&pb.Upvotes{Total: upvotes}); err != nil {
			log.Printf("send error %v", err)
			return err
		}
		log.Printf("sending data: %v", upvotes)
	}
}
