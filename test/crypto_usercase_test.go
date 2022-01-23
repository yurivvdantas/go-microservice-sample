package test

import (
	"context"
	pb "go-microservice-sample/api"
	"log"
	"net"
	"testing"

	usecases "go-microservice-sample/internal/crypto-votes-service/usescases"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/emptypb"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

// Test grpc call to get a crypto by id.
func TestGrpcGetById(t *testing.T) {
	req := &pb.CryptoId{Id: 1}
	grpcClient, conn := getCryptoClientGrpc()
	defer conn.Close()
	res, err := grpcClient.Get(context.Background(), req)
	if err != nil || res == nil {
		t.Fatalf("Test fail on get crypto by ID: %d with err: %v", req.Id, err)
	}
}

// Test grpc call to get all crypto
func TestGrpcGetAll(t *testing.T) {
	grpcClient, conn := getCryptoClientGrpc()
	defer conn.Close()
	res, err := grpcClient.GetAll(context.Background(), &emptypb.Empty{})
	if err != nil || res == nil {
		t.Fatalf("Test fail on get all crypto with err: %v", err)
	}
}

// Set up a connection to the server.
func getCryptoClientGrpc() (pb.CryptosClient, *grpc.ClientConn) {
	//conn, err := grpc.Dial(configs.Crypto_votes_service_address,
	//grpc.WithInsecure(), grpc.WithBlock())
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return pb.NewCryptosClient(conn), conn
}

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterCryptosServer(s, &usecases.CryptoServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}
