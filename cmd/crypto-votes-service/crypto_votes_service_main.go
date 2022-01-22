package main

import (
	"log"
	"net"

	pb "go-microservice-sample/api"
	"go-microservice-sample/configs"
	usecases "go-microservice-sample/internal/crypto-votes-service/usescases"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", configs.Crypto_votes_service_address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCryptosServer(s, &usecases.CryptoServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
