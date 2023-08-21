// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Genre struct {
	ID   uuid.UUID
	User uuid.UUID
	Name string
}

type Movie struct {
	ID         uuid.UUID
	Title      string
	Rating     string
	Cast       []string
	Director   string
	Poster     pgtype.Text
	User       uuid.UUID
	UserRating pgtype.Int4
}

type MovieGenre struct {
	ID    uuid.UUID
	Movie uuid.UUID
	Genre uuid.UUID
}

type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	CreatedAt pgtype.Timestamp
	DeletedAt pgtype.Timestamp
	LastLogin pgtype.Timestamp
}

type UserMovie struct {
	ID    uuid.UUID
	User  uuid.UUID
	Movie uuid.UUID
}
