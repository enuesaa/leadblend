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
SELECT * FROM islands WHERE id = ?;
-- name: CreateIsland :one
INSERT INTO islands (title, content, comment) VALUES (?, ?, ?) RETURNING *;
-- name: DeleteIsland :exec
DELETE FROM islands WHERE id = ?;


-- name: ListPatterns :many
SELECT * FROM patterns;
-- name: GetPattern :one
SELECT * FROM patterns WHERE id = ?;
-- name: CreatePattern :one
INSERT INTO patterns (title, priority) VALUES (?, ?) RETURNING *;
-- name: DeletePattern :exec
DELETE FROM patterns WHERE id = ?;



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
INSERT INTO stones (pattern_id, island_id, data) VALUES (?, ?, ?) RETURNING *;
-- name: DeleteStone :exec
DELETE FROM stones WHERE id = ?;



-- name: ListResources :many
SELECT * FROM resources;
-- name: GetResource :one
SELECT * FROM resources WHERE id = ?;
-- name: CreateResource :one
INSERT INTO resources (id, marker) VALUES (?, ?) RETURNING *;
-- name: DeleteResource :exec
DELETE FROM resources WHERE id = ?;



-- name: ListTags :many
SELECT * FROM tags;
-- name: GetTag :one
SELECT * FROM tags WHERE id = ?;
-- name: CreateTag :one
INSERT INTO tags (resource_id, key, value) VALUES (?, ?, ?) RETURNING *;
-- name: DeleteTag :exec
DELETE FROM tags WHERE id = ?;



-- name: ListHistories :many
SELECT * FROM histories;
-- name: GetHistorys :one
SELECT * FROM histories WHERE id = ?;
-- name: CreateHistory :one
INSERT INTO histories (resource_id, comment) VALUES (?, ?) RETURNING *;
-- name: DeleteHistory :exec
DELETE FROM histories WHERE id = ?;
