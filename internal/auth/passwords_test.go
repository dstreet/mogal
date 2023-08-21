package auth_test

import (
	"testing"

	"github.com/dstreet/mogal/internal/auth"
	"github.com/stretchr/testify/assert"
)

func Test_BcryptPasswordHasher(t *testing.T) {
	hasher := auth.NewBcryptPasswordHasher(12)
	password := "supersecret"

	hash, err := hasher.HashPassword(password)
	assert.NoError(t, err)

	verified := hasher.Verify(password, hash)
	assert.True(t, verified)
}
