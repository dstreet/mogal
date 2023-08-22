package db

import (
	"context"
	"log/slog"

	"github.com/dstreet/mogal/internal/genre"
	"github.com/gofrs/uuid"
)

type DBGenreRepository struct {
	logger *slog.Logger

	queries *Queries
}

func NewDBGenreRepository(logger *slog.Logger, db DBTX) *DBGenreRepository {
	return &DBGenreRepository{
		logger:  logger,
		queries: New(db),
	}
}

func (repo *DBGenreRepository) GetAllForUser(ctx context.Context, userID string) ([]genre.Genre, error) {
	repo.logger.Info("getting genres for user", "user", userID)

	userUUID, err := uuid.FromString(userID)
	if err != nil {
		repo.logger.Error("invalid uuid")
		return nil, err
	}

	genres, err := repo.queries.GetUserGenres(ctx, userUUID)
	if err != nil {
		repo.logger.Error("failed to get genres from DB", "err", err)
		return nil, err
	}

	userGenres := make([]genre.Genre, len(genres))
	for i, g := range genres {
		userGenres[i] = genre.Genre{
			ID:   g.ID.String(),
			Name: g.Name,
		}
	}

	return userGenres, nil
}

func (repo *DBGenreRepository) CreateGenreForUser(ctx context.Context, genreInput genre.GenreInput, userID string) (genre.Genre, error) {
	repo.logger.Info("create a new genre for user", "user", userID)

	userUUID, err := uuid.FromString(userID)
	if err != nil {
		repo.logger.Error("invalid uuid")
		return genre.Genre{}, err
	}

	args := CreateGenreForUserParams{
		Name: genreInput.Name,
		User: userUUID,
	}

	res, err := repo.queries.CreateGenreForUser(ctx, args)
	if err != nil {
		repo.logger.Error("failed to create genre for user", "err", err)
		return genre.Genre{}, err
	}

	return genre.Genre{
		ID:   res.ID.String(),
		Name: res.Name,
	}, nil
}
