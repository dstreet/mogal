package auth

import (
	"time"

	"github.com/dstreet/mogal/internal/user"
	"github.com/golang-jwt/jwt/v5"
)

type JWTProvider struct {
	issuer     string
	signingKey string
}

func NewJWTProvider(issuer string, signingKey string) *JWTProvider {
	return &JWTProvider{
		issuer:     issuer,
		signingKey: signingKey,
	}
}

func (p *JWTProvider) CreateToken(u user.User, expires time.Duration) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    p.issuer,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expires)),
		Subject:   u.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(p.signingKey))
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (p *JWTProvider) VerifyToken(token string) bool {
	return true
}
