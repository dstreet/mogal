package http

import (
	"time"

	"github.com/dstreet/mogal/internal/user"
)

//go:generate mockery --name TokenProvider
type TokenProvider interface {
	CreateToken(u user.User) (string, error)
	VerifyToken(token string) (userID string, err error)
	TokenDuration() time.Duration
}
