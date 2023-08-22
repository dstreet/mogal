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

-- name: GetMoviesForUser :many
SELECT * FROM movies
WHERE "user" = $1
ORDER BY "title" ASC;

-- name: GetMoviesForUserAndGenre :many
SELECT movies.* FROM movies
INNER JOIN movie_genres mg on mg.movie = movies.id
WHERE movies.user = $1
AND mg.genre = $2
ORDER BY movies.title ASC;

-- name: GetUserMovie :one
SELECT * FROM movies
WHERE "id" = $1
AND "user" = $2;