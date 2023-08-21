-- name: CreateMovie :one
INSERT INTO movies (
  "title",
  "rating",
  "cast",
  "director",
  "poster",
  "user_rating",
  "user"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: AddMovieGenres :copyfrom
INSERT INTO movie_genres (
  "movie",
  "genre"
) VALUES (
  $1, $2
);

-- name: GetMovieGenres :many
SELECT genres.* FROM genres
INNER JOIN movie_genres mg ON mg.genre = genres.id
WHERE mg.movie = $1
ORDER BY genres.name ASC;