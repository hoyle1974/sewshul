package main

import (
	"context"
	"log"

	"github.com/hoyle1974/sewshul/microservice"
	pb "github.org/hoyle1974/sewshul/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLoginServiceServer
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Received: %v/%v", in.GetUsername(), in.GetPassword())
	return &pb.LoginResponse{Message: "Hello " + in.GetUsername()}, nil

}

func main() {
	microservice.Start("login", func(s *grpc.Server) {
		pb.RegisterLoginServiceServer(s, &server{})
	})
}
