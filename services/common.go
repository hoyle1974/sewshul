package services

import (
	"database/sql"

	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

type AppCtx struct {
	s   *grpc.Server
	log zerolog.Logger
	db  *sql.DB
}

func NewAppCtx(s *grpc.Server, log zerolog.Logger, db *sql.DB) AppCtx {
	return AppCtx{
		s:   s,
		log: log,
		db:  db,
	}
}
