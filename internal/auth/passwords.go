package auth

import "golang.org/x/crypto/bcrypt"

type BcryptPasswordHasher struct {
	rounds int
}

func NewBcryptPasswordHasher(rounds int) *BcryptPasswordHasher {
	return &BcryptPasswordHasher{
		rounds: rounds,
	}
}

func (ph *BcryptPasswordHasher) HashPassword(p string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(p), ph.rounds)
	if err != nil {
		return "", err
	}

	return string(h), nil
}

func (ph *BcryptPasswordHasher) Verify(p string, h string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	return err == nil
}
