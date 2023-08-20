package user

import "context"

type UserRepository interface {
	// Create a user with the provided credentials.
	Create(ctx context.Context, email string, password string) (User, error)

	// Delete a user with the provided ID.
	Delete(ctx context.Context, ID string) (User, error)

	// Verify that the user exists with the provided credentials and perform a
	// login operation.
	Login(ctx context.Context, email string, password string) (User, error)
}
