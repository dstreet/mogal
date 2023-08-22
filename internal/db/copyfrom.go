// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: copyfrom.go

package db

import (
	"context"
)

// iteratorForAddMovieGenres implements pgx.CopyFromSource.
type iteratorForAddMovieGenres struct {
	rows                 []AddMovieGenresParams
	skippedFirstNextCall bool
}

func (r *iteratorForAddMovieGenres) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForAddMovieGenres) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].Movie,
		r.rows[0].Genre,
	}, nil
}

func (r iteratorForAddMovieGenres) Err() error {
	return nil
}

func (q *Queries) AddMovieGenres(ctx context.Context, arg []AddMovieGenresParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"movie_genres"}, []string{"movie", "genre"}, &iteratorForAddMovieGenres{rows: arg})
}

// iteratorForCreateGenresForUser implements pgx.CopyFromSource.
type iteratorForCreateGenresForUser struct {
	rows                 []CreateGenresForUserParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreateGenresForUser) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateGenresForUser) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].Name,
		r.rows[0].User,
	}, nil
}

func (r iteratorForCreateGenresForUser) Err() error {
	return nil
}

func (q *Queries) CreateGenresForUser(ctx context.Context, arg []CreateGenresForUserParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"genres"}, []string{"name", "user"}, &iteratorForCreateGenresForUser{rows: arg})
}
