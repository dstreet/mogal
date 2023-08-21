package genre

import "context"

//go:generate mockery --name GenreRepository
type GenreRepository interface {
	// Get all the genres for a user.
	GetAllForUser(ctx context.Context, userID string) ([]Genre, error)
}
