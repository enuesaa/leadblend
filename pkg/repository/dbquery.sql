-- name: ListPlanets :many
SELECT * FROM planets;
-- name: GetPlanet :one
SELECT * FROM planets WHERE id = ?;
-- name: GetPlanetByName :one
SELECT * FROM planets WHERE name = ?;
-- name: CreatePlanet :one
INSERT INTO planets (id, name, comment) VALUES (?, ?, ?) RETURNING *;
-- name: UpdatePlanetName :exec
UPDATE planets SET name = sqlc.arg(new_name) WHERE name = sqlc.arg(name);
-- name: DeletePlanet :exec
DELETE FROM planets WHERE name = ?;



-- name: ListIslands :many
SELECT * FROM islands;
-- name: ListIslandsByTitle :many
SELECT * FROM islands where title = ?;
-- name: ListIslandsByPlanetId :many
SELECT * FROM islands where planet_id = ?;
-- name: GetIsland :one
SELECT * FROM islands WHERE id = ?;
-- name: CreateIsland :one
INSERT INTO islands (id, planet_id, title, content, comment) VALUES (?, ?, ?, ?, ?) RETURNING *;
-- name: DeleteIsland :exec
DELETE FROM islands WHERE id = ?;


-- name: ListPatterns :many
SELECT * FROM patterns;
-- name: GetPattern :one
SELECT * FROM patterns WHERE id = ?;
-- name: CreatePattern :one
INSERT INTO patterns (id, title, priority, color) VALUES (?, ?, ?, ?) RETURNING *;
-- name: DeletePattern :exec
DELETE FROM patterns WHERE id = ?;



-- name: ListTraits :many
SELECT * FROM traits;
-- name: ListTraitsByPatternIds :many
SELECT * FROM traits WHERE pattern_id IN (sqlc.slice('pattern_ids'));
-- name: GetTrait :one
SELECT * FROM traits WHERE id = ?;
-- name: CreateTrait :one
INSERT INTO traits (id, pattern_id, path, type, default_value, required) VALUES (?, ?, ?, ?, ?, ?) RETURNING *;
-- name: DeleteTrait :exec
DELETE FROM traits WHERE id = ?;



-- name: ListStones :many
SELECT * FROM stones;
-- name: ListStonesByIslandId :many
SELECT * FROM stones WHERE island_id = ?;
-- name: ListComets :many
SELECT * FROM stones WHERE island_id IS NULL;
-- name: GetStone :one
SELECT * FROM stones WHERE id = ?;
-- name: CreateStone :one
INSERT INTO stones (id, pattern_id, island_id, data) VALUES (?, ?, ?, ?) RETURNING *;
-- name: UpdateStone :exec
UPDATE stones SET island_id = ? WHERE id = ?;
-- name: UpdateStonePatternId :exec
UPDATE stones SET pattern_id = ? WHERE id = ?;
-- name: DeleteStone :exec
DELETE FROM stones WHERE id = ?;



-- name: ListTags :many
SELECT * FROM tags;
-- name: GetTag :one
SELECT * FROM tags WHERE id = ?;
-- name: CreateTag :one
INSERT INTO tags (resource, key, value) VALUES (?, ?, ?) RETURNING *;
-- name: DeleteTag :exec
DELETE FROM tags WHERE id = ?;



-- name: ListHistories :many
SELECT * FROM histories;
-- name: GetHistorys :one
SELECT * FROM histories WHERE id = ?;
-- name: CreateHistory :one
INSERT INTO histories (resource, comment) VALUES (?, ?) RETURNING *;
-- name: DeleteHistory :exec
DELETE FROM histories WHERE id = ?;
