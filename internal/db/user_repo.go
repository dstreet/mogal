package db

import (
	"context"
	"errors"
	"log/slog"

	"github.com/dstreet/mogal/internal/user"
	"github.com/gofrs/uuid"
)

type DBUserRepository struct {
	logger *slog.Logger

	queries        *Queries
	passwordHasher PasswordHasher
}

type PasswordHasher interface {
	HashPassword(p string) (string, error)
	Verify(p string, h string) bool
}

func NewDBUserRepository(logger *slog.Logger, db DBTX, ph PasswordHasher) *DBUserRepository {
	return &DBUserRepository{
		logger:         logger,
		queries:        New(db),
		passwordHasher: ph,
	}
}

// Create a new user with the given email and password
func (repo *DBUserRepository) Create(ctx context.Context, email string, password string) (user.User, error) {
	var userObj user.User
	repo.logger.Info("creating new user", "email", email)

	if password == "" {
		repo.logger.Warn("invalid password received")
		return userObj, user.ErrInvalidPassword
	}

	hash, err := repo.passwordHasher.HashPassword(password)
	if err != nil {
		repo.logger.Error("failed to hash password", "err", err)
		return userObj, errors.New("failed to hash password")
	}

	params := CreateUserParams{
		Email:    email,
		Password: hash,
	}

	u, err := repo.queries.CreateUser(ctx, params)
	if err != nil {
		repo.logger.Error("failed to create user", "err", err)
		return userObj, err
	}

	userObj = updateUserObjectWithDBUser(userObj, u)

	repo.logger.Info("successfully created user", "email", email, "id", userObj.ID)

	return userObj, nil
}

// Delete a user given their ID
func (repo *DBUserRepository) Delete(ctx context.Context, ID string) (user.User, error) {
	var userObj user.User
	repo.logger.Info("deleting user", "ID", ID)

	uuidID, err := uuid.FromString(ID)
	if err != nil {
		repo.logger.Error("invalid id when deleting user", "id", ID)
		return userObj, user.ErrInvalidID
	}

	u, err := repo.queries.DeleteUser(ctx, uuidID)
	if err != nil {
		repo.logger.Error("failed to delete user from DB", "err", err)
		return userObj, err
	}

	userObj = updateUserObjectWithDBUser(userObj, u)

	repo.logger.Info("successfully deleted user", "ID", ID)
	return userObj, nil
}

// Perform a login operation for a user given email and password
func (repo *DBUserRepository) Login(ctx context.Context, email string, password string) (user.User, error) {
	var userObj user.User
	repo.logger.Info("attempting to login user", "email", email)

	u, err := repo.queries.GetUserWithEmail(ctx, email)
	if err != nil {
		repo.logger.Error("failed to get user", "email", email)
		return userObj, &user.NotFoundError{Email: &email}
	}

	if !repo.passwordHasher.Verify(password, u.Password) {
		repo.logger.Error("failed to login user due to incorrect password")
		return userObj, user.ErrIncorrectPassword
	}

	err = repo.queries.UpdateLastLogin(ctx, u.ID)
	if err != nil {
		repo.logger.Error("failed to update last login date for user", "id", u.ID.String())
		return userObj, &user.LoginFailedError{ID: u.ID.String()}
	}

	userObj = updateUserObjectWithDBUser(userObj, u)

	return userObj, nil
}

func updateUserObjectWithDBUser(userObj user.User, dbUser User) user.User {
	userObj.ID = dbUser.ID.String()
	userObj.Email = dbUser.Email
	userObj.CreatedAt = dbUser.CreatedAt
	userObj.LastLogin = dbUser.LastLogin

	userObj.Active = !dbUser.DeletedAt.Valid

	if !userObj.Active {
		userObj.DeletedAt = &dbUser.DeletedAt.Time
	}

	return userObj
}
