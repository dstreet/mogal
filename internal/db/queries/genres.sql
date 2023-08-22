-- name: GetUserGenres :many
SELECT * FROM genres
WHERE "user" = $1
ORDER BY "name" ASC;

-- name: GetUserGenresByName :many
SELECT * FROM genres
WHERE "name" = ANY(sqlc.arg(names)::text[])
AND "user" = $1;

-- name: CreateGenresForUser :copyfrom
INSERT INTO genres (
  "name",
  "user"
) VALUES (
  $1, $2
);
