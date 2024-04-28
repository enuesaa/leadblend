-- name: ListPlanets :many
SELECT * FROM planets;

-- name: GetPlanet :one
SELECT * FROM planets
WHERE name = ?;

-- name: CreatePlanet :one
INSERT INTO planets (
  name, comment
) VALUES (
  ?, ?
)
RETURNING *;

-- name: DeletePlanet :exec
DELETE FROM planets
WHERE name = ?;


-- name: ListIslands :many
SELECT * FROM islands;

-- name: GetIsland :one
SELECT * FROM islands
WHERE name = ?;

-- name: CreateIsland :one
INSERT INTO islands (
  name, content, comment
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: DeleteIsland :exec
DELETE FROM islands
WHERE name = ?;


-- name: ListIslandTags :many
SELECT * FROM island_tags;

-- name: CreateIslandTag :one
INSERT INTO island_tags (
  island_id, key, value
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: DeleteIslandTag :exec
DELETE FROM island_tags
WHERE id = ?;


