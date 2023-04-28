package microservice

import (
	"flag"
	"fmt"
	"net"
	"os"

	grpczerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 8443, "The server port")
	ip   = flag.String("ip", "0.0.0.0", "address")
	env  = flag.String("env", "dev", "environment")
)

func Start(name string, register func(*grpc.Server)) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if *env == "dev" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	log.Info().Msg("Starting " + name + " microservice")

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *ip, *port))
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("failed to listen")
	}

	s := grpc.NewServer(
		middleware.WithUnaryServerChain(
			logging.UnaryServerInterceptor(grpczerolog.InterceptorLogger(log.Logger)),
		),
		middleware.WithStreamServerChain(
			logging.StreamServerInterceptor(grpczerolog.InterceptorLogger(log.Logger)),
		),
	)
	register(s)
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal().AnErr("error", err).Msg("failed to serve")
	}

}
