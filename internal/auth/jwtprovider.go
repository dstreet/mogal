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

// Generate a new authorization token for the user.
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

// Verify that the authorization token is valid.
func (p *JWTProvider) VerifyToken(token string) (string, error) {
	var userID string

	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return userID, ErrUnexpectedSigningTyp
		}

		return []byte(p.signingKey), nil
	})

	if err != nil {
		return userID, err
	}

	issuer, err := t.Claims.GetIssuer()
	if err != nil {
		return userID, err
	}

	if issuer != p.issuer {
		return userID, ErrMismatchedIssuer
	}

	userID, err = t.Claims.GetSubject()
	if err != nil {
		return userID, err
	}

	return userID, err
}
