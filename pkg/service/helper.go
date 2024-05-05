package service

import (
	"context"
	"database/sql"
)

func ctx() context.Context {
	return context.Background()
}

func nstr(value string) sql.NullString {
	return sql.NullString{String: value, Valid: true}
}
