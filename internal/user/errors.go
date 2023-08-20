package user

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidPassword   = errors.New("invalid password")
	ErrInvalidID         = errors.New("invalid ID")
	ErrIncorrectPassword = errors.New("incorrect password")
)

type NotFoundError struct {
	ID    *string
	Email *string
}

func (e *NotFoundError) Error() string {
	if e.ID != nil {
		return fmt.Sprintf("user with ID %s not found", *e.ID)
	}

	if e.Email != nil {
		return fmt.Sprintf("user with email %s not found", *e.Email)
	}

	return "user not found"
}

type LoginFailedError struct {
	ID string
}

func (e *LoginFailedError) Error() string {
	return fmt.Sprintf("failed to login user with ID %s", e.ID)
}
