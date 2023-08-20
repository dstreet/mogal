package auth

import "errors"

var (
	ErrMismatchedIssuer     = errors.New("token issuers do not match")
	ErrUnexpectedSigningTyp = errors.New("unexpected signing type for token")
)
