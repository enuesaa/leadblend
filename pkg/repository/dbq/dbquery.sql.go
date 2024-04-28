// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: dbquery.sql

package dbq

import (
	"context"
	"database/sql"
)

const createHistory = `-- name: CreateHistory :one
INSERT INTO histories (resource_id, comment) VALUES (?, ?) RETURNING id, resource_id, comment, created, updated
`

type CreateHistoryParams struct {
	ResourceID string
	Comment    string
}

func (q *Queries) CreateHistory(ctx context.Context, arg CreateHistoryParams) (History, error) {
	row := q.db.QueryRowContext(ctx, createHistory, arg.ResourceID, arg.Comment)
	var i History
	err := row.Scan(
		&i.ID,
		&i.ResourceID,
		&i.Comment,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const createIsland = `-- name: CreateIsland :one
INSERT INTO islands (title, content, comment) VALUES (?, ?, ?) RETURNING id, title, content, comment, created, updated
`

type CreateIslandParams struct {
	Title   string
	Content string
	Comment string
}

func (q *Queries) CreateIsland(ctx context.Context, arg CreateIslandParams) (Island, error) {
	row := q.db.QueryRowContext(ctx, createIsland, arg.Title, arg.Content, arg.Comment)
	var i Island
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.Comment,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const createPattern = `-- name: CreatePattern :one
INSERT INTO patterns (title, priority) VALUES (?, ?) RETURNING id, title, priority, created, updated
`

type CreatePatternParams struct {
	Title    string
	Priority sql.NullInt64
}

func (q *Queries) CreatePattern(ctx context.Context, arg CreatePatternParams) (Pattern, error) {
	row := q.db.QueryRowContext(ctx, createPattern, arg.Title, arg.Priority)
	var i Pattern
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Priority,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const createPlanet = `-- name: CreatePlanet :one
INSERT INTO planets (name, comment) VALUES (?, ?) RETURNING id, name, comment, created, updated
`

type CreatePlanetParams struct {
	Name    string
	Comment string
}

func (q *Queries) CreatePlanet(ctx context.Context, arg CreatePlanetParams) (Planet, error) {
	row := q.db.QueryRowContext(ctx, createPlanet, arg.Name, arg.Comment)
	var i Planet
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Comment,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const createResource = `-- name: CreateResource :one
INSERT INTO resources (id, marker) VALUES (?, ?) RETURNING id, marker
`

type CreateResourceParams struct {
	ID     string
	Marker string
}

func (q *Queries) CreateResource(ctx context.Context, arg CreateResourceParams) (Resource, error) {
	row := q.db.QueryRowContext(ctx, createResource, arg.ID, arg.Marker)
	var i Resource
	err := row.Scan(&i.ID, &i.Marker)
	return i, err
}

const createStone = `-- name: CreateStone :one
INSERT INTO stones (pattern_id, island_id, data) VALUES (?, ?, ?) RETURNING id, pattern_id, island_id, data, created, updated
`

type CreateStoneParams struct {
	PatternID sql.NullInt64
	IslandID  sql.NullInt64
	Data      string
}

func (q *Queries) CreateStone(ctx context.Context, arg CreateStoneParams) (Stone, error) {
	row := q.db.QueryRowContext(ctx, createStone, arg.PatternID, arg.IslandID, arg.Data)
	var i Stone
	err := row.Scan(
		&i.ID,
		&i.PatternID,
		&i.IslandID,
		&i.Data,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const createTag = `-- name: CreateTag :one
INSERT INTO tags (resource_id, key, value) VALUES (?, ?, ?) RETURNING id, resource_id, "key", value, created, updated
`

type CreateTagParams struct {
	ResourceID string
	Key        string
	Value      string
}

func (q *Queries) CreateTag(ctx context.Context, arg CreateTagParams) (Tag, error) {
	row := q.db.QueryRowContext(ctx, createTag, arg.ResourceID, arg.Key, arg.Value)
	var i Tag
	err := row.Scan(
		&i.ID,
		&i.ResourceID,
		&i.Key,
		&i.Value,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const createTrait = `-- name: CreateTrait :one
INSERT INTO traits (pattern_id, path, type, default_value, required) VALUES (?, ?, ?, ?, ?) RETURNING id, pattern_id, path, type, default_value, required, created, updated
`

type CreateTraitParams struct {
	PatternID    int64
	Path         string
	Type         string
	DefaultValue string
	Required     bool
}

func (q *Queries) CreateTrait(ctx context.Context, arg CreateTraitParams) (Trait, error) {
	row := q.db.QueryRowContext(ctx, createTrait,
		arg.PatternID,
		arg.Path,
		arg.Type,
		arg.DefaultValue,
		arg.Required,
	)
	var i Trait
	err := row.Scan(
		&i.ID,
		&i.PatternID,
		&i.Path,
		&i.Type,
		&i.DefaultValue,
		&i.Required,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const deleteHistory = `-- name: DeleteHistory :exec
DELETE FROM histories WHERE id = ?
`

func (q *Queries) DeleteHistory(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteHistory, id)
	return err
}

const deleteIsland = `-- name: DeleteIsland :exec
DELETE FROM islands WHERE id = ?
`

func (q *Queries) DeleteIsland(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteIsland, id)
	return err
}

const deletePattern = `-- name: DeletePattern :exec
DELETE FROM patterns WHERE id = ?
`

func (q *Queries) DeletePattern(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePattern, id)
	return err
}

const deletePlanet = `-- name: DeletePlanet :exec
DELETE FROM planets WHERE name = ?
`

func (q *Queries) DeletePlanet(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, deletePlanet, name)
	return err
}

const deleteResource = `-- name: DeleteResource :exec
DELETE FROM resources WHERE id = ?
`

func (q *Queries) DeleteResource(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteResource, id)
	return err
}

const deleteStone = `-- name: DeleteStone :exec
DELETE FROM stones WHERE id = ?
`

func (q *Queries) DeleteStone(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteStone, id)
	return err
}

const deleteTag = `-- name: DeleteTag :exec
DELETE FROM tags WHERE id = ?
`

func (q *Queries) DeleteTag(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTag, id)
	return err
}

const deleteTrait = `-- name: DeleteTrait :exec
DELETE FROM traits WHERE id = ?
`

func (q *Queries) DeleteTrait(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTrait, id)
	return err
}

const getHistorys = `-- name: GetHistorys :one
SELECT id, resource_id, comment, created, updated FROM histories WHERE id = ?
`

func (q *Queries) GetHistorys(ctx context.Context, id int64) (History, error) {
	row := q.db.QueryRowContext(ctx, getHistorys, id)
	var i History
	err := row.Scan(
		&i.ID,
		&i.ResourceID,
		&i.Comment,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getIsland = `-- name: GetIsland :one
SELECT id, title, content, comment, created, updated FROM islands WHERE id = ?
`

func (q *Queries) GetIsland(ctx context.Context, id int64) (Island, error) {
	row := q.db.QueryRowContext(ctx, getIsland, id)
	var i Island
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.Comment,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getPattern = `-- name: GetPattern :one
SELECT id, title, priority, created, updated FROM patterns WHERE id = ?
`

func (q *Queries) GetPattern(ctx context.Context, id int64) (Pattern, error) {
	row := q.db.QueryRowContext(ctx, getPattern, id)
	var i Pattern
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Priority,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getPlanet = `-- name: GetPlanet :one
SELECT id, name, comment, created, updated FROM planets WHERE name = ?
`

func (q *Queries) GetPlanet(ctx context.Context, name string) (Planet, error) {
	row := q.db.QueryRowContext(ctx, getPlanet, name)
	var i Planet
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Comment,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getResource = `-- name: GetResource :one
SELECT id, marker FROM resources WHERE id = ?
`

func (q *Queries) GetResource(ctx context.Context, id string) (Resource, error) {
	row := q.db.QueryRowContext(ctx, getResource, id)
	var i Resource
	err := row.Scan(&i.ID, &i.Marker)
	return i, err
}

const getStone = `-- name: GetStone :one
SELECT id, pattern_id, island_id, data, created, updated FROM stones WHERE id = ?
`

func (q *Queries) GetStone(ctx context.Context, id int64) (Stone, error) {
	row := q.db.QueryRowContext(ctx, getStone, id)
	var i Stone
	err := row.Scan(
		&i.ID,
		&i.PatternID,
		&i.IslandID,
		&i.Data,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getTag = `-- name: GetTag :one
SELECT id, resource_id, "key", value, created, updated FROM tags WHERE id = ?
`

func (q *Queries) GetTag(ctx context.Context, id int64) (Tag, error) {
	row := q.db.QueryRowContext(ctx, getTag, id)
	var i Tag
	err := row.Scan(
		&i.ID,
		&i.ResourceID,
		&i.Key,
		&i.Value,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getTrait = `-- name: GetTrait :one
SELECT id, pattern_id, path, type, default_value, required, created, updated FROM traits WHERE id = ?
`

func (q *Queries) GetTrait(ctx context.Context, id int64) (Trait, error) {
	row := q.db.QueryRowContext(ctx, getTrait, id)
	var i Trait
	err := row.Scan(
		&i.ID,
		&i.PatternID,
		&i.Path,
		&i.Type,
		&i.DefaultValue,
		&i.Required,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listHistories = `-- name: ListHistories :many
SELECT id, resource_id, comment, created, updated FROM histories
`

func (q *Queries) ListHistories(ctx context.Context) ([]History, error) {
	rows, err := q.db.QueryContext(ctx, listHistories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []History
	for rows.Next() {
		var i History
		if err := rows.Scan(
			&i.ID,
			&i.ResourceID,
			&i.Comment,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listIslands = `-- name: ListIslands :many
SELECT id, title, content, comment, created, updated FROM islands
`

func (q *Queries) ListIslands(ctx context.Context) ([]Island, error) {
	rows, err := q.db.QueryContext(ctx, listIslands)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Island
	for rows.Next() {
		var i Island
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.Comment,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPatterns = `-- name: ListPatterns :many
SELECT id, title, priority, created, updated FROM patterns
`

func (q *Queries) ListPatterns(ctx context.Context) ([]Pattern, error) {
	rows, err := q.db.QueryContext(ctx, listPatterns)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Pattern
	for rows.Next() {
		var i Pattern
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Priority,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPlanets = `-- name: ListPlanets :many
SELECT id, name, comment, created, updated FROM planets
`

func (q *Queries) ListPlanets(ctx context.Context) ([]Planet, error) {
	rows, err := q.db.QueryContext(ctx, listPlanets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Planet
	for rows.Next() {
		var i Planet
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Comment,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listResources = `-- name: ListResources :many
SELECT id, marker FROM resources
`

func (q *Queries) ListResources(ctx context.Context) ([]Resource, error) {
	rows, err := q.db.QueryContext(ctx, listResources)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Resource
	for rows.Next() {
		var i Resource
		if err := rows.Scan(&i.ID, &i.Marker); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listStones = `-- name: ListStones :many
SELECT id, pattern_id, island_id, data, created, updated FROM stones
`

func (q *Queries) ListStones(ctx context.Context) ([]Stone, error) {
	rows, err := q.db.QueryContext(ctx, listStones)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Stone
	for rows.Next() {
		var i Stone
		if err := rows.Scan(
			&i.ID,
			&i.PatternID,
			&i.IslandID,
			&i.Data,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTags = `-- name: ListTags :many
SELECT id, resource_id, "key", value, created, updated FROM tags
`

func (q *Queries) ListTags(ctx context.Context) ([]Tag, error) {
	rows, err := q.db.QueryContext(ctx, listTags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tag
	for rows.Next() {
		var i Tag
		if err := rows.Scan(
			&i.ID,
			&i.ResourceID,
			&i.Key,
			&i.Value,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTraits = `-- name: ListTraits :many
SELECT id, pattern_id, path, type, default_value, required, created, updated FROM traits
`

func (q *Queries) ListTraits(ctx context.Context) ([]Trait, error) {
	rows, err := q.db.QueryContext(ctx, listTraits)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Trait
	for rows.Next() {
		var i Trait
		if err := rows.Scan(
			&i.ID,
			&i.PatternID,
			&i.Path,
			&i.Type,
			&i.DefaultValue,
			&i.Required,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
