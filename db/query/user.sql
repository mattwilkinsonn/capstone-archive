-- name: GetUserById :one
SELECT
    *
FROM
    users
WHERE
    id = $1
LIMIT 1;

-- name: GetUserByUsername :one
SELECT
    *
FROM
    users
WHERE
    username = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT
    *
FROM
    users
WHERE
    email = $1
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, username, email, PASSWORD)
    VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUserRole :one
UPDATE
    users
SET
    ROLE = $1
WHERE
    id = $2
RETURNING
    *;

