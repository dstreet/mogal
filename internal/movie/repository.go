package movie

import (
	"context"

	"github.com/dstreet/mogal/internal/genre"
)

//go:generate mockery --name MovieRepository
type MovieRepository interface {
	// Create a new movie for a user.
	CreateMovie(ctx context.Context, movie MovieInput, userID string) (Movie, error)

	// Get the genres for a movie.
	GetGenres(ctx context.Context, movieID string) ([]genre.Genre, error)
}
