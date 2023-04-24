package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	pb "github.org/hoyle1974/sewshul/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
	ip   = flag.String("ip", "0.0.0.0", "address")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedAccountServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	log.Printf("Received: %v/%v", in.GetUsername(), in.GetPassword())
	return &pb.CreateAccountResponse{Message: "Hello " + in.GetUsername(), AccountId: "id"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *ip, *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAccountServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
