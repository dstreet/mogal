package graphql_test

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/dstreet/mogal/internal/genre"
	"github.com/dstreet/mogal/internal/graphql"
	"github.com/dstreet/mogal/internal/graphql/model"
	"github.com/dstreet/mogal/internal/http"
	"github.com/dstreet/mogal/internal/movie"
	"github.com/dstreet/mogal/internal/user"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var testLogger = slog.New(slog.NewTextHandler(os.Stderr, nil))

func Test_MutationResolver_Login(t *testing.T) {
	tp := http.NewMockTokenProvider(t)
	ur := user.NewMockUserRepository(t)

	resolver := &graphql.Resolver{
		Logger:         testLogger,
		UserRepository: ur,
		TokenProvider:  tp,
	}

	mr := resolver.Mutation()

	t.Run("throw unauthorized error when user does not exist", func(t *testing.T) {
		input := model.LoginInput{
			Email:    "test@testerton.com",
			Password: "supersecret",
		}

		ur.EXPECT().Login(mock.Anything, input.Email, input.Password).
			Return(user.User{}, &user.NotFoundError{}).
			Once()

		_, err := mr.Login(context.Background(), input)
		assert.ErrorIs(t, err, http.ErrUnauthorized)
	})

	t.Run("throw unauthorized error when password is incorrect", func(t *testing.T) {
		input := model.LoginInput{
			Email:    "test@testerton.com",
			Password: "supersecret",
		}

		ur.EXPECT().Login(mock.Anything, input.Email, input.Password).
			Return(user.User{}, user.ErrIncorrectPassword).
			Once()

		_, err := mr.Login(context.Background(), input)
		assert.ErrorIs(t, err, http.ErrUnauthorized)
	})

	t.Run("returns authorization when successfully authenticated", func(t *testing.T) {
		input := model.LoginInput{
			Email:    "test@testerton.com",
			Password: "supersecret",
		}

		authUser := user.User{}
		token := "faketoken"
		expiresIn := 900

		ur.EXPECT().Login(mock.Anything, input.Email, input.Password).
			Return(user.User{}, nil).
			Once()

		tp.EXPECT().CreateToken(authUser).
			Return(token, nil).
			Once()

		tp.EXPECT().TokenDuration().
			Return(time.Second * time.Duration(expiresIn)).
			Once()

		authorization, err := mr.Login(context.Background(), input)
		assert.NoError(t, err)

		assert.Equal(t, token, authorization.Token)
		assert.Equal(t, expiresIn, authorization.ExpiresIn)
	})
}

func Test_MutationResolver_Register(t *testing.T) {
	tp := http.NewMockTokenProvider(t)
	ur := user.NewMockUserRepository(t)

	resolver := &graphql.Resolver{
		Logger:         testLogger,
		UserRepository: ur,
		TokenProvider:  tp,
	}

	mr := resolver.Mutation()

	t.Run("returns error when failed to create user", func(t *testing.T) {
		input := model.RegisterInput{
			Email:    "test@testerton.com",
			Password: "supersecret",
		}

		ur.EXPECT().Create(mock.Anything, input.Email, input.Password).
			Return(user.User{}, errors.New("uh oh")).
			Once()

		_, err := mr.Register(context.Background(), input)
		assert.Error(t, err)
	})

	t.Run("returns error when failed to create token", func(t *testing.T) {
		input := model.RegisterInput{
			Email:    "test@testerton.com",
			Password: "supersecret",
		}

		authUser := user.User{}

		ur.EXPECT().Create(mock.Anything, input.Email, input.Password).
			Return(authUser, nil).
			Once()

		tp.EXPECT().CreateToken(authUser).
			Return("", errors.New("uh oh")).
			Once()

		_, err := mr.Register(context.Background(), input)
		assert.Error(t, err)
	})

	t.Run("returns authorization when user was successfully created", func(t *testing.T) {
		input := model.RegisterInput{
			Email:    "test@testerton.com",
			Password: "supersecret",
		}

		authUser := user.User{}
		token := "faketoken"
		expiresIn := 900

		ur.EXPECT().Create(mock.Anything, input.Email, input.Password).
			Return(authUser, nil).
			Once()

		tp.EXPECT().CreateToken(authUser).
			Return(token, nil).
			Once()

		tp.EXPECT().TokenDuration().
			Return(time.Second * time.Duration(expiresIn)).
			Once()

		authorization, err := mr.Register(context.Background(), input)
		assert.NoError(t, err)

		assert.Equal(t, token, authorization.Token)
		assert.Equal(t, expiresIn, authorization.ExpiresIn)
	})
}

func Test_QueryResolver_ListGenres(t *testing.T) {
	gr := genre.NewMockGenreRepository(t)

	resolver := &graphql.Resolver{
		Logger:          testLogger,
		GenreRepository: gr,
	}

	qr := resolver.Query()

	t.Run("throw unauthorized when user isn't authenticated", func(t *testing.T) {
		_, err := qr.ListGenres(context.Background())
		assert.ErrorIs(t, err, http.ErrUnauthorized)
	})

	t.Run("returns user genres", func(t *testing.T) {
		user := user.User{ID: "1234567890"}
		ctx := context.WithValue(context.Background(), http.UserCtxKey, &user)

		genres := []genre.Genre{
			{
				ID:   "1",
				Name: "Action",
			},
			{
				ID:   "2",
				Name: "Comedy",
			},
			{
				ID:   "1",
				Name: "Horror",
			},
		}

		gr.EXPECT().GetAllForUser(mock.Anything, user.ID).
			Return(genres, nil).
			Once()

		res, err := qr.ListGenres(ctx)
		assert.NoError(t, err)

		assert.Len(t, res, 3)
		for i, g := range genres {
			assert.Equal(t, g.ID, res[i].ID)
			assert.Equal(t, g.Name, res[i].Name)
		}
	})
}

func Test_MutationResolver_CreateMovie(t *testing.T) {
	tp := http.NewMockTokenProvider(t)
	movieRepo := movie.NewMockMovieRepository(t)
	fieldCollector := graphql.NewMockFieldCollector(t)

	resolver := &graphql.Resolver{
		Logger:          testLogger,
		MovieRepository: movieRepo,
		TokenProvider:   tp,
		FieldCollector:  fieldCollector,
	}

	mr := resolver.Mutation()

	t.Run("throws unathorized when the user isn't authenticated", func(t *testing.T) {
		_, err := mr.CreateMovie(context.Background(), model.CreateMovieInput{})
		assert.ErrorIs(t, err, http.ErrUnauthorized)
	})

	t.Run("creates and returns the movie", func(t *testing.T) {
		user := user.User{ID: "1234567890"}
		ctx := context.WithValue(context.Background(), http.UserCtxKey, &user)

		input := model.CreateMovieInput{
			Title:    "Super cool movie",
			Rating:   "PG-13",
			Cast:     []string{"Tom Waits"},
			Director: "Jim Jarmusch",
		}

		fieldCollector.EXPECT().CollectAllFields(mock.Anything).
			Return([]string{}).
			Once()

		movieRepo.EXPECT().CreateMovie(mock.Anything, movie.MovieInput{
			Title:    input.Title,
			Rating:   input.Rating,
			Cast:     input.Cast,
			Director: input.Director,
		}, user.ID).
			Return(movie.Movie{
				ID:       "1234567890",
				Title:    input.Title,
				Rating:   input.Rating,
				Cast:     input.Cast,
				Director: input.Director,
			}, nil).
			Once()

		res, err := mr.CreateMovie(ctx, input)
		assert.NoError(t, err)

		assert.Equal(t, input.Title, res.Title)
		assert.Equal(t, input.Rating, res.Rating)
		assert.Equal(t, input.Cast, res.Cast)
		assert.Equal(t, input.Director, res.Director)
	})

	t.Run("creates a new movie with genres and returns the genres", func(t *testing.T) {
		user := user.User{ID: "1234567890"}
		ctx := context.WithValue(context.Background(), http.UserCtxKey, &user)

		input := model.CreateMovieInput{
			Title:    "Super cool movie",
			Rating:   "PG-13",
			Cast:     []string{"Tom Waits"},
			Director: "Jim Jarmusch",
			Genres:   []string{"123", "456"},
		}

		movieID := "1234567890"

		genres := []genre.Genre{
			{
				ID:   "123",
				Name: "Action",
			},
			{
				ID:   "456",
				Name: "Horror",
			},
		}

		fieldCollector.EXPECT().CollectAllFields(mock.Anything).
			Return([]string{"genres"}).
			Once()

		movieRepo.EXPECT().CreateMovie(mock.Anything, movie.MovieInput{
			Title:    input.Title,
			Rating:   input.Rating,
			Cast:     input.Cast,
			Director: input.Director,
			Genres:   input.Genres,
		}, user.ID).
			Return(movie.Movie{
				ID:       movieID,
				Title:    input.Title,
				Rating:   input.Rating,
				Cast:     input.Cast,
				Director: input.Director,
			}, nil).
			Once()

		movieRepo.EXPECT().GetGenres(mock.Anything, movieID).
			Return(genres, nil).
			Once()

		res, err := mr.CreateMovie(ctx, input)
		assert.NoError(t, err)

		assert.Equal(t, input.Title, res.Title)
		assert.Equal(t, input.Rating, res.Rating)
		assert.Equal(t, input.Cast, res.Cast)
		assert.Equal(t, input.Director, res.Director)

		assert.Len(t, res.Genres, 2)
		for i, g := range genres {
			assert.Equal(t, g.ID, res.Genres[i].ID)
			assert.Equal(t, g.Name, res.Genres[i].Name)
		}
	})
}
