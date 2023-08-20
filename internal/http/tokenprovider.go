package http

import (
	"time"

	"github.com/dstreet/mogal/internal/user"
)

type TokenProvider interface {
	CreateToken(u user.User, expires time.Duration) (string, error)
	VerifyToken(token string) (userID string, err error)
}
