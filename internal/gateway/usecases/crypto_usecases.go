package usecases

import (
	"context"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	pb "go-microservice-sample/api"
	"go-microservice-sample/configs"
	"go-microservice-sample/internal/gateway/model"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func GetCryptoByID(ctxRequest *gin.Context) {
	id, err := strconv.ParseInt(ctxRequest.Param("id"), 10, 64)
	if err != nil {
		log.Fatalf("The id must be a int: %v", err)
	}
	grpcClient, conn := getCryptoClientGrpc()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := grpcClient.Get(ctx, &pb.CryptoId{Id: id})
	if err != nil {
		log.Printf("something get wrong on get user: %v", err)
		ctxRequest.IndentedJSON(http.StatusBadRequest,
			model.HTTPError{Cause: err.Error(), Detail: "something get wrong on get user", Status: 400})
	}
	log.Printf("Result: %s", r.Name)

	ctxRequest.IndentedJSON(http.StatusOK, r)
}

func GetAllCrypto(ctxRequest *gin.Context) {
	log.Printf("Getting all cryptos")
	grpcClient, conn := getCryptoClientGrpc()
	defer conn.Close()

	stream, err := grpcClient.GetAll(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	var cryptos []pb.Crypto
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("cannot receive %v", err)
		}
		cryptos = append(cryptos, pb.Crypto{Id: resp.Id,
			Name:        resp.Name,
			Code:        resp.Code,
			Upvote:      resp.Upvote,
			Downvote:    resp.Downvote,
			Description: resp.Description})
		log.Printf("Resp received: %s", resp.Name)
	}
	ctxRequest.IndentedJSON(http.StatusOK, cryptos)
}

func AddCrypto(ctxRequest *gin.Context) {
	var newCrypto model.Crypto_request

	if err := ctxRequest.BindJSON(&newCrypto); err != nil {
		ctxRequest.IndentedJSON(http.StatusBadRequest,
			model.HTTPError{Cause: err.Error(), Detail: "it's not a valid crypto", Status: 400})
		return
	}

	grpcClient, conn := getCryptoClientGrpc()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := grpcClient.AddCrypto(ctx, &pb.Crypto{
		Name:        newCrypto.Name,
		Code:        newCrypto.Code,
		Description: newCrypto.Description})
	if err != nil {
		log.Printf("something get wrong on add user: %v", err)
		ctxRequest.IndentedJSON(http.StatusBadRequest,
			model.HTTPError{Cause: err.Error(), Detail: "something get wrong on add user", Status: 500})
		return
	}
	log.Printf("The crypto was inserted and the id is: %d", r.Id)

	ctxRequest.IndentedJSON(http.StatusOK, r)
}

func UpvoteCrypto(ctxRequest *gin.Context) {
	id, err := strconv.ParseInt(ctxRequest.Param("id"), 10, 64)
	if err != nil {
		log.Fatalf("The id must be a int: %v", err)
	}
	grpcClient, conn := getCryptoClientGrpc()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = grpcClient.Upvote(ctx, &pb.CryptoId{Id: id})
	if err != nil {
		log.Printf("something get wrong on upvote: %v", err)
		ctxRequest.IndentedJSON(http.StatusBadRequest,
			model.HTTPError{Cause: err.Error(), Detail: "something get wrong on upvote", Status: 400})
		return
	}
	log.Printf("Upvote finish")

	ctxRequest.IndentedJSON(http.StatusNoContent, nil)
}

// Set up a connection to the server.
func getCryptoClientGrpc() (pb.CryptosClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(configs.Crypto_votes_service_address,
		grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return pb.NewCryptosClient(conn), conn
}
