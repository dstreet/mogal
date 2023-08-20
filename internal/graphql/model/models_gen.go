// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Authorization struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}