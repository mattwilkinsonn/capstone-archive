-- name: GetCapstoneById :one
SELECT
    *
FROM
    capstones
WHERE
    id = $1
LIMIT 1;

-- name: GetCapstonesWithCursor :many
SELECT
    *
FROM
    capstones
WHERE
    created_at < $1
ORDER BY
    created_at DESC
LIMIT $2;

-- name: GetCapstones :many
SELECT
    *
FROM
    capstones
ORDER BY
    created_at DESC
LIMIT $1;

-- name: SearchCapstones :many
SELECT
    *
FROM
    capstones c
WHERE
    to_tsvector(c.Title) || to_tsvector(c.Description) || to_tsvector(c.Author) || to_tsvector(c.Semester) @@ to_tsquery('english', $1)
LIMIT $2 OFFSET $3;

-- SELECT
--     *
-- FROM (
--     SELECT
--         to_tsvector(c.Title) || to_tsvector(c.Description) || to_tsvector(c.Author) || to_tsvector(c.Semester) AS document,
--         *
--     FROM
--         capstones c) search
-- WHERE
--     search.document @@ to_tsquery('english', $1)
-- LIMIT $2 OFFSET $3;
-- name: CreateCapstone :one
INSERT INTO capstones (id, created_at, updated_at, title, description, author, semester)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING
    *;

