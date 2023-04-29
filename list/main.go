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

func (s *server) GetSocialList(ctx context.Context, in *pb.GetSocialListRequest) (*pb.GetSocialListResponse, error) {
	log.Printf("Received: %v", in.GetUserId())

	stmt := `select id, entity_id from "lists" where "owner_id" = $1 and "list_type" = $2`
	rows, err := s.db.Query(stmt, in.GetUserId(), in.GetListType().String())
	if err != nil {
		return &pb.GetSocialListResponse{Error: microservice.ErrToProto(err)}, err
	}

	entities := make([]string, 0)
	defer rows.Close()
	for rows.Next() {
		var id, entity_id string
		rows.Scan(&id, &entity_id)
		entities = append(entities, entity_id)
	}

	return &pb.GetSocialListResponse{Ids: entities}, nil
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
