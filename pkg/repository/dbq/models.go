// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package dbq

import (
	"database/sql"
)

type History struct {
	ID       string
	Resource string
	Comment  string
	Created  sql.NullTime
	Updated  sql.NullTime
}

type Island struct {
	ID       string
	PlanetID string
	Title    string
	Content  string
	Comment  string
	Created  sql.NullTime
	Updated  sql.NullTime
}

type Pattern struct {
	ID       string
	Title    string
	Priority sql.NullInt64
	Created  sql.NullTime
	Updated  sql.NullTime
}

type Planet struct {
	ID      string
	Name    string
	Comment string
	Created sql.NullTime
	Updated sql.NullTime
}

type Stone struct {
	ID        string
	PatternID sql.NullString
	IslandID  sql.NullString
	Data      string
	Created   sql.NullTime
	Updated   sql.NullTime
}

type Tag struct {
	ID       string
	Resource string
	Key      string
	Value    string
	Created  sql.NullTime
	Updated  sql.NullTime
}

type Trait struct {
	ID           string
	PatternID    int64
	Path         string
	Type         string
	DefaultValue string
	Required     bool
	Created      sql.NullTime
	Updated      sql.NullTime
}
