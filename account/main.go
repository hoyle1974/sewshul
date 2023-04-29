package main

import (
	"context"

	"github.com/hoyle1974/sewshul/microservice"
	pb "github.com/hoyle1974/sewshul/proto"
	"github.com/hoyle1974/sewshul/services"
	_ "github.com/rs/zerolog/log"
)

type server struct {
	pb.UnimplementedAccountServiceServer
	appCtx services.AppCtx
}

func (s *server) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {

	id, err := services.CreateAccount(s.app, in.GetUsername(), in.GetPassword())
	if err != nil {
		return &pb.CreateAccountResponse{Error: microservice.ErrToProto(err)}, err
	}

	return &pb.CreateAccountResponse{Message: "Hello " + in.GetUsername(), AccountId: id.String()}, nil
}

func main() {
	microservice.Start("account", func(appCtx services.AppCtx) {
		pb.RegisterAccountServiceServer(s, &server{appCtx: appCtx})
	})
}
