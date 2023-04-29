package main

import (
	"context"

	"github.com/hoyle1974/sewshul/microservice"
	pb "github.com/hoyle1974/sewshul/proto"
	"github.com/hoyle1974/sewshul/services"
)

type server struct {
	pb.UnimplementedLoginServiceServer
	appCtx services.AppCtx
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {

	accountId, err := services.Login(s.appCtx, in.GetUsername(), in.GetPassword())
	if err != nil {
		return &pb.LoginResponse{Error: microservice.ErrToProto(err)}, err

	}

	return &pb.LoginResponse{Message: "Hello " + in.GetUsername(), UserId: accountId.String()}, nil

}

func register(appCtx services.AppCtx) {
	pb.RegisterLoginServiceServer(appCtx.Server, &server{appCtx: appCtx})
}

func main() {
	microservice.Start("login", register)
}
