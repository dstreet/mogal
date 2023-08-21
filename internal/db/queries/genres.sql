-- name: GetUserGenres :many
SELECT * from genres
WHERE "user" = $1
ORDER BY "name" ASC;