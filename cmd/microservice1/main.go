package main

import (
	"context"
	"log"
	"net"

	pb "go-microservice-sample/api"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedUsersServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Get(ctx context.Context, in *pb.UserId) (*pb.User, error) {
	log.Printf("Received userId: %v", in.Id)
	return &pb.User{Id: in.Id, Name: "New user"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUsersServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
