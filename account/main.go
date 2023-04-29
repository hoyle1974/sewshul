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
	pb.UnimplementedAccountServiceServer
	db *sql.DB
}

func (s *server) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	log.Printf("Received: %v/%v", in.GetUsername(), in.GetPassword())

	hash, err := microservice.HashPassword(in.GetPassword())
	if err != nil {
		return &pb.CreateAccountResponse{Error: microservice.ErrToProto(err)}, err
	}

	stmt := `insert into "users"("id", "username","password_hash") values(gen_random_uuid(),$1, $2) returning id`
	row := s.db.QueryRow(stmt, in.GetUsername(), hash)
	if row.Err() != nil {
		return &pb.CreateAccountResponse{Error: microservice.ErrToProto(row.Err())}, err
	}

	var id string
	row.Scan(&id)

	return &pb.CreateAccountResponse{Message: "Hello " + in.GetUsername(), AccountId: id}, nil
}

func main() {
	microservice.Start("account", func(s *grpc.Server, db *sql.DB) {
		pb.RegisterAccountServiceServer(s, &server{db: db})
	})
}
