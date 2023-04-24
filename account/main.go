package main

import (
	"context"
	"log"

	"github.com/hoyle1974/sewshul/microservice"
	pb "github.com/hoyle1974/sewshul/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAccountServiceServer
}

func (s *server) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	log.Printf("Received: %v/%v", in.GetUsername(), in.GetPassword())
	return &pb.CreateAccountResponse{Message: "Hello " + in.GetUsername(), AccountId: "id"}, nil
}

func main() {
	microservice.Start("account", func(s *grpc.Server) {
		pb.RegisterAccountServiceServer(s, &server{})
	})
}
