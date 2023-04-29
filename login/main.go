package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/hoyle1974/sewshul/microservice"
	pb "github.com/hoyle1974/sewshul/proto"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLoginServiceServer
	db *sql.DB
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Received: %v/%v", in.GetUsername(), in.GetPassword())

	stmt := `select id, password_hash from "users" where "username"= $1`
	row := s.db.QueryRow(stmt, in.GetUsername())

	if row.Err() != nil {
		return &pb.LoginResponse{Error: microservice.ErrToProto(row.Err())}, row.Err()
	}

	var id, hash string
	row.Scan(&id, &hash)

	var err error
	if err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(in.GetPassword())); err != nil {
		return &pb.LoginResponse{Message: "Error" + in.GetUsername()}, err
	}

	return &pb.LoginResponse{Message: "Hello " + in.GetUsername(), UserId: id}, nil

}

func main() {
	microservice.Start("login", func(s *grpc.Server, db *sql.DB) {
		pb.RegisterLoginServiceServer(s, &server{db: db})
	})
}
