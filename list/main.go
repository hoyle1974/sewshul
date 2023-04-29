package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/hoyle1974/sewshul/microservice"
	pb "github.com/hoyle1974/sewshul/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSocialListServiceServer
	db *sql.DB
}

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
	microservice.Start("sociallist", func(s *grpc.Server, db *sql.DB) {
		pb.RegisterSocialListServiceServer(s, &server{db: db})
	})
}
