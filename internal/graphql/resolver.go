package graphql

import (
	"log/slog"
	"time"

	"github.com/dstreet/mogal/internal/user"
)

type TokenProvider interface {
	CreateToken(u user.User, expires time.Duration) (string, error)
	VerifyToken(token string) bool
}

type Resolver struct {
	Logger *slog.Logger

	UserRepository user.UserRepository
	TokenProvider  TokenProvider
}
