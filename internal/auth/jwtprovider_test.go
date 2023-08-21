package auth_test

import (
	"testing"
	"time"

	"github.com/dstreet/mogal/internal/auth"
	"github.com/dstreet/mogal/internal/user"
	"github.com/stretchr/testify/assert"
)

func Test_JWTProvider(t *testing.T) {
	t.Run("creates valid token than can be verified", func(t *testing.T) {
		tp := auth.NewJWTProvider("testissuer", "keyboardcat", time.Second*900)
		authUser := user.User{ID: "1234567890"}

		token, err := tp.CreateToken(authUser)
		assert.NoError(t, err)

		ID, err := tp.VerifyToken(token)
		assert.NoError(t, err)
		assert.Equal(t, authUser.ID, ID)
	})

	t.Run("fails to verify when issuers don't match", func(t *testing.T) {
		tpIssuer := auth.NewJWTProvider("testissuer1", "keyboardcat", time.Second*900)
		tpVerifier := auth.NewJWTProvider("testissuer2", "keyboardcat", time.Second*900)

		authUser := user.User{ID: "1234567890"}

		token, err := tpIssuer.CreateToken(authUser)
		assert.NoError(t, err)

		_, err = tpVerifier.VerifyToken(token)
		assert.Error(t, err)
	})
}
