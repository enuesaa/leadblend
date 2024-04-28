-- name: ListSpaces :many
SELECT * FROM spaces;

-- name: GetSpace :one
SELECT * FROM spaces
WHERE name = ?;

-- name: CreateSpace :one
INSERT INTO spaces (
  name
) VALUES (
  ?
)
RETURNING *;

-- name: DeleteSpace :exec
DELETE FROM spaces
WHERE name = ?;
