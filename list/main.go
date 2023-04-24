package main

import (
	"context"
	"log"

	"github.com/hoyle1974/sewshul/microservice"
	pb "github.com/hoyle1974/sewshul/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSocialListServiceServer
}

// func (s *server) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
// 	log.Printf("Received: %v/%v", in.GetUsername(), in.GetPassword())
// 	return &pb.CreateAccountResponse{Message: "Hello " + in.GetUsername(), AccountId: "id"}, nil
// }

func (s *server) GetSocialList(ctx context.Context, in *pb.SocialListRequest) (*pb.SocialListResponse, error) {
	log.Printf("Received: %v", in.GetUserId())
	return &pb.SocialListResponse{}, nil

}

func (s *server) AddToSocialList(ctx context.Context, in *pb.AddToSocialListRequest) (*pb.AddToSocialListResponse, error) {
	return &pb.AddToSocialListResponse{}, nil
}

func (s *server) RemoveFromSocialList(ctx context.Context, in *pb.RemoveFromSocialListRequest) (*pb.RemoveFromSocialListResponse, error) {
	return &pb.RemoveFromSocialListResponse{}, nil
}

func main() {
	microservice.Start("sociallist", func(s *grpc.Server) {
		pb.RegisterSocialListServiceServer(s, &server{})
	})
}
