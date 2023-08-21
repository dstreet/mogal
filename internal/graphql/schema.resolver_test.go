package graphql_test

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/dstreet/mogal/internal/graphql"
	"github.com/dstreet/mogal/internal/graphql/model"
	"github.com/dstreet/mogal/internal/http"
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
