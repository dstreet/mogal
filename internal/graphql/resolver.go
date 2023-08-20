package graphql

import (
	"log/slog"

	"github.com/dstreet/mogal/internal/http"
	"github.com/dstreet/mogal/internal/user"
)

type Resolver struct {
	Logger *slog.Logger

	UserRepository user.UserRepository
	TokenProvider  http.TokenProvider
}
