-- name: GetUserGenres :many
SELECT * from genres
WHERE "user" = $1
ORDER BY "name" ASC;

-- name: CreateGenreForUser :one
INSERT INTO genres (
  "name",
  "user"
) VALUES (
  $1, $2
)
RETURNING *;
