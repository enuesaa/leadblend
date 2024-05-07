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

func nint64(value int64) sql.NullInt64 {
	return sql.NullInt64{Int64: value, Valid: true}
}
