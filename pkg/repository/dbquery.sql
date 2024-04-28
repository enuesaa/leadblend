-- name: ListPlanets :many
SELECT * FROM planets;
-- name: GetPlanet :one
SELECT * FROM planets WHERE name = ?;
-- name: CreatePlanet :one
INSERT INTO planets (name, comment) VALUES (?, ?) RETURNING *;
-- name: DeletePlanet :exec
DELETE FROM planets WHERE name = ?;



-- name: ListIslands :many
SELECT * FROM islands;
-- name: GetIsland :one
SELECT * FROM islands WHERE name = ?;
-- name: CreateIsland :one
INSERT INTO islands (name, content, comment) VALUES (?, ?, ?) RETURNING *;
-- name: DeleteIsland :exec
DELETE FROM islands WHERE name = ?;



-- name: ListIslandTags :many
SELECT * FROM island_tags;
-- name: CreateIslandTag :one
INSERT INTO island_tags (island_id, key, value) VALUES (?, ?, ?) RETURNING *;
-- name: DeleteIslandTag :exec
DELETE FROM island_tags WHERE id = ?;



-- name: ListPatterns :many
SELECT * FROM patterns;
-- name: GetPattern :one
SELECT * FROM patterns WHERE name = ?;
-- name: CreatePattern :one
INSERT INTO patterns (name, priority) VALUES (?, ?) RETURNING *;
-- name: DeletePattern :exec
DELETE FROM patterns WHERE name = ?;



-- name: ListTraits :many
SELECT * FROM traits;
-- name: GetTrait :one
SELECT * FROM traits WHERE id = ?;
-- name: CreateTrait :one
INSERT INTO traits (pattern_id, path, type, default_value, required) VALUES (?, ?, ?, ?, ?) RETURNING *;
-- name: DeleteTrait :exec
DELETE FROM traits WHERE id = ?;



-- name: ListStones :many
SELECT * FROM stones;
-- name: GetStone :one
SELECT * FROM stones WHERE id = ?;
-- name: CreateStone :one
INSERT INTO stones (id, pattern_id, island_id, data) VALUES (?, ?, ?, ?) RETURNING *;
-- name: DeleteStone :exec
DELETE FROM stones WHERE id = ?;



-- name: ListHistories :many
SELECT * FROM histories;
-- name: CreateHistory :one
INSERT INTO histories (id, resource, comment) VALUES (?, ?, ?) RETURNING *;
-- name: DeleteHistory :exec
DELETE FROM histories WHERE id = ?;
