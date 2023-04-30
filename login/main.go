package main

import (
	"context"
	"net"

	"github.com/hoyle1974/sewshul/microservice"
	pb "github.com/hoyle1974/sewshul/proto"
	"github.com/hoyle1974/sewshul/services"
)

type server struct {
	pb.UnimplementedLoginServiceServer
	appCtx services.AppCtx
}

func UserContactsToPB(contacts []services.UserContact) []*pb.UserContact {
	out := make([]*pb.UserContact, len(contacts))
	for _, contact := range contacts {
		c := pb.UserContact{
			UserId: contact.AccountID.String(),
			ClientAddress: &pb.ClientAddress{
				IpAddress: contact.Ip.String(),
				Port:      contact.Port,
			},
		}
		out = append(out, &c)
	}

	return out
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {

	accountId, err := services.Login(
		s.appCtx,
		in.GetUsername(),
		in.GetPassword(),
		net.ParseIP(in.GetClientAddress().GetIpAddress()),
		in.GetClientAddress().GetPort(),
	)
	if err != nil {
		return &pb.LoginResponse{Error: microservice.ErrToProto(err)}, err
	}

	followsIds, err := services.GetSocialList(s.appCtx, accountId, services.SocialListType_FOLLOWS)
	if err != nil {
		return &pb.LoginResponse{Error: microservice.ErrToProto(err)}, err
	}
	followingIds, err := services.GetSocialList(s.appCtx, accountId, services.SocialListType_FOLLOWING)
	if err != nil {
		return &pb.LoginResponse{Error: microservice.ErrToProto(err)}, err
	}
	blockedIds, err := services.GetSocialList(s.appCtx, accountId, services.SocialListType_BLOCKED)
	if err != nil {
		return &pb.LoginResponse{Error: microservice.ErrToProto(err)}, err
	}

	follows, err := services.GetUserContacts(s.appCtx, followsIds)
	if err != nil {
		return &pb.LoginResponse{Error: microservice.ErrToProto(err)}, err
	}
	following, err := services.GetUserContacts(s.appCtx, followingIds)
	if err != nil {
		return &pb.LoginResponse{Error: microservice.ErrToProto(err)}, err
	}

	return &pb.LoginResponse{
		Message:   "Hello " + in.GetUsername(),
		UserId:    accountId.String(),
		Blocked:   services.AccountIdsToStrings(blockedIds),
		Follows:   UserContactsToPB(follows),
		Following: UserContactsToPB(following),
	}, nil

}

func register(appCtx services.AppCtx) {
	pb.RegisterLoginServiceServer(appCtx.Server, &server{appCtx: appCtx})
}

func main() {
	microservice.Start("login", register)
}
