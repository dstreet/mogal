package db

import (
	"context"
	"log/slog"

	"github.com/dstreet/mogal/internal/genre"
	"github.com/dstreet/mogal/internal/movie"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type DBMovieRepository struct {
	logger *slog.Logger

	queries *Queries
}

func NewDBMovieRepository(logger *slog.Logger, db DBTX) *DBMovieRepository {
	return &DBMovieRepository{
		logger:  logger,
		queries: New(db),
	}
}

func (repo *DBMovieRepository) CreateMovie(ctx context.Context, input movie.MovieInput, userID string) (movie.Movie, error) {
	repo.logger.Info("creating new movie for user", "user", userID)

	userUUID, err := uuid.FromString(userID)
	if err != nil {
		repo.logger.Error("invalid uuid")
		return movie.Movie{}, err
	}

	dbInput := CreateMovieParams{
		Title:    input.Title,
		Rating:   input.Rating,
		Cast:     input.Cast,
		Director: input.Director,
		User:     userUUID,
	}

	if input.Poster != nil {
		dbInput.Poster = pgtype.Text{String: *input.Poster, Valid: true}
	}

	if input.UserRating != nil {
		dbInput.UserRating = pgtype.Int4{Int32: *input.UserRating, Valid: true}
	}

	dbMovie, err := repo.queries.CreateMovie(ctx, dbInput)
	if err != nil {
		repo.logger.Error("failed to create movie", "err", err)
		return movie.Movie{}, nil
	}

	if len(input.Genres) > 0 {
		movieGenres := make([]AddMovieGenresParams, len(input.Genres))
		for i, g := range input.Genres {
			genreUUID, err := uuid.FromString(g)
			if err != nil {
				return movie.Movie{}, err
			}

			movieGenres[i] = AddMovieGenresParams{
				Movie: dbMovie.ID,
				Genre: genreUUID,
			}
		}

		_, err := repo.queries.AddMovieGenres(ctx, movieGenres)
		if err != nil {
			repo.logger.Error("failed to add generes when creating movie", "movie", dbMovie.ID, "err", err)
			return movie.Movie{}, err
		}
	}

	var poster *string

	if dbMovie.Poster.Valid {
		poster = &dbMovie.Poster.String
	}

	var userRating *int32

	if dbMovie.UserRating.Valid {
		userRating = &dbMovie.UserRating.Int32
	}

	return movie.Movie{
		ID:         dbMovie.ID.String(),
		Title:      dbMovie.Title,
		Rating:     dbMovie.Rating,
		Cast:       dbMovie.Cast,
		Director:   dbMovie.Director,
		Poster:     poster,
		UserRating: userRating,
	}, nil
}

func (repo *DBMovieRepository) GetGenres(ctx context.Context, movieID string) ([]genre.Genre, error) {
	repo.logger.Info("getting genres for movie", "movie", movieID)

	movieUUID, err := uuid.FromString(movieID)
	if err != nil {
		repo.logger.Error("invalid uuid")
		return nil, err
	}

	dbRes, err := repo.queries.GetMovieGenres(ctx, movieUUID)
	if err != nil {
		return nil, err
	}

	genres := make([]genre.Genre, len(dbRes))
	for i, g := range dbRes {
		genres[i] = genre.Genre{
			ID:   g.ID.String(),
			Name: g.Name,
		}
	}

	return genres, nil
}

func (repo *DBMovieRepository) GetMoviesForUser(ctx context.Context, userID string, genre *string) ([]movie.Movie, error) {
	repo.logger.Info("getting movies for user", "user", userID, "genre", genre)

	userUUID, err := uuid.FromString(userID)
	if err != nil {
		repo.logger.Error("invalid uuid for user")
		return nil, err
	}

	var movieRes []Movie

	if genre != nil {
		genreUUID, err := uuid.FromString(*genre)
		if err != nil {
			repo.logger.Error("invalid uuid for genre")
			return nil, err
		}

		movieRes, err = repo.queries.GetMoviesForUserAndGenre(ctx, GetMoviesForUserAndGenreParams{
			User:  userUUID,
			Genre: genreUUID,
		})

		if err != nil {
			repo.logger.Error("failed to get movie for user and genre", "err", err)
			return nil, err
		}
	} else {
		movieRes, err = repo.queries.GetMoviesForUser(ctx, userUUID)
		if err != nil {
			repo.logger.Error("failed to get movie for user", "err", err)
			return nil, err
		}
	}

	movies := make([]movie.Movie, len(movieRes))
	for i, m := range movieRes {
		movies[i] = updateMovieObjectWithDBMovie(movie.Movie{}, m)
	}

	return movies, nil
}

func (repo *DBMovieRepository) GetMovieForUser(ctx context.Context, userID string, movieId string) (movie.Movie, error) {
	repo.logger.Info("getting movie for user", "user", userID, "movie", movieId)

	var movieObj movie.Movie

	userUUID, err := uuid.FromString(userID)
	if err != nil {
		repo.logger.Error("invalid uuid for user")
		return movieObj, err
	}

	movieUUID, err := uuid.FromString(movieId)
	if err != nil {
		repo.logger.Error("invalid uuid for movie")
		return movieObj, err
	}

	res, err := repo.queries.GetUserMovie(ctx, GetUserMovieParams{
		User: userUUID,
		ID:   movieUUID,
	})
	if err != nil {
		repo.logger.Error("failed to get movie", "err", err)
		return movie.Movie{}, err
	}

	updateMovieObjectWithDBMovie(movieObj, res)

	return movieObj, nil
}

func updateMovieObjectWithDBMovie(movieObj movie.Movie, dbMovie Movie) movie.Movie {
	movieObj.ID = dbMovie.ID.String()
	movieObj.Title = dbMovie.Title
	movieObj.Rating = dbMovie.Rating
	movieObj.Cast = dbMovie.Cast
	movieObj.Director = dbMovie.Director

	if dbMovie.Poster.Valid {
		movieObj.Poster = &dbMovie.Poster.String
	}

	if dbMovie.UserRating.Valid {
		movieObj.UserRating = &dbMovie.UserRating.Int32
	}

	return movieObj
}
