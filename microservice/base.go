package microservice

import (
	"flag"
	"fmt"
	"net"
	"os"

	grpczerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	pb "github.com/hoyle1974/sewshul/proto"
	"github.com/hoyle1974/sewshul/services"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func ErrToProto(err error) *pb.Error {
	return &pb.Error{Msg: err.Error()}
}

func Start(name string, register func(appCtx services.AppCtx)) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if *env == "dev" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	log.Info().Msg("Connecting to database")
	db, err := DB()
	if err != nil {
		log.Fatal().AnErr("Error connecting to database", err).Msg("error")
		return
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

	appCtx := services.AppCtx{
		s:   s,
		log: log,
		db:  db,
	}

	register(appCtx)
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal().AnErr("error", err).Msg("failed to serve")
	}

}
