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

func (repo *DBGenreRepository) CreateGenresForUser(ctx context.Context, genreInput []genre.GenreInput, userID string) ([]genre.Genre, error) {
	repo.logger.Info("create a new genre for user", "user", userID)

	userUUID, err := uuid.FromString(userID)
	if err != nil {
		repo.logger.Error("invalid uuid")
		return nil, err
	}

	args := make([]CreateGenresForUserParams, len(genreInput))
	for i, gi := range genreInput {
		args[i] = CreateGenresForUserParams{
			Name: gi.Name,
			User: userUUID,
		}
	}

	_, err = repo.queries.CreateGenresForUser(ctx, args)
	if err != nil {
		repo.logger.Error("failed to create genre for user", "err", err)
		return nil, err
	}

	names := make([]string, len(genreInput))
	for i, gi := range genreInput {
		names[i] = gi.Name
	}

	res, err := repo.queries.GetUserGenresByName(ctx, GetUserGenresByNameParams{
		User:  userUUID,
		Names: names,
	})
	if err != nil {
		repo.logger.Error("failed to get newly creted genres", "err", err)
		return nil, err
	}

	genres := make([]genre.Genre, len(res))
	for i, r := range res {
		genres[i] = genre.Genre{
			ID:   r.ID.String(),
			Name: r.Name,
		}
	}

	return genres, nil
}
