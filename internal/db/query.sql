-- name: GetUserWithEmail :one
SELECT * FROM users
WHERE email = $1
LIMIT 1;

-- name: GetUserWithId :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

-- name: DeleteUser :one
UPDATE users
SET deleted_at = NOW()
WHERE id = $1
RETURNING *;

-- name: CreateUser :one
INSERT INTO users (
  email,
  password,
  created_at,
  last_login
) VALUES (
  $1,
  $2,
  NOW(),
  NOW()
)
RETURNING *;

-- name: UpdateLastLogin :exec
UPDATE users
SET last_login = NOW()
WHERE id = $1;