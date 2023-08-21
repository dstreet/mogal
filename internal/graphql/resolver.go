package graphql

import (
	"context"
	"log/slog"

	"github.com/dstreet/mogal/internal/genre"
	"github.com/dstreet/mogal/internal/http"
	"github.com/dstreet/mogal/internal/movie"
	"github.com/dstreet/mogal/internal/user"
)

//go:generate mockery --name FieldCollector
type FieldCollector interface {
	CollectAllFields(context.Context) []string
}

type Resolver struct {
	Logger *slog.Logger

	UserRepository  user.UserRepository
	GenreRepository genre.GenreRepository
	MovieRepository movie.MovieRepository
	TokenProvider   http.TokenProvider

	FieldCollector FieldCollector
}
