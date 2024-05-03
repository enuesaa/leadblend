package repository

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	_ "modernc.org/sqlite"
	"os"

	"github.com/enuesaa/leadblend/pkg/repository/dbq"
)

type DBRepositoryInterface interface {
	IsDBExist() bool
	Migrate() error
	Open() error
	Close() error
	Query() *dbq.Queries
}

//go:embed dbschema.sql
var dbMigrateQuery string

type DBRepository struct {
	db *sql.DB
}

func (repo *DBRepository) dbpath() string {
	return ".leadblend/main.leadblend/data.db"
}

func (repo *DBRepository) dsn(path string) string {
	return fmt.Sprintf("file:%s?_fk=1", path)
}

func (repo *DBRepository) IsDBExist() bool {
	if _, err := os.Stat(repo.dbpath()); os.IsNotExist(err) {
		return false
	}
	return true
}

func (repo *DBRepository) Migrate() error {
	db, err := sql.Open("sqlite", repo.dsn(repo.dbpath()))
	if err != nil {
		return err
	}
	ctx := context.Background()
	if _, err := db.ExecContext(ctx, "SELECT * FROM teas"); err != nil {
		if _, err := db.ExecContext(ctx, dbMigrateQuery); err != nil {
			return err
		}
	}
	return nil
}

func (repo *DBRepository) Open() error {
	if repo.db != nil {
		return nil
	}
	db, err := sql.Open("sqlite", repo.dsn(repo.dbpath()))
	if err != nil {
		return err
	}
	repo.db = db
	return nil
}

func (repo *DBRepository) Close() error {
	if repo.db == nil {
		return nil
	}
	if err := repo.db.Close(); err != nil {
		return err
	}
	repo.db = nil
	return nil
}

func (repo *DBRepository) Query() *dbq.Queries {
	return dbq.New(repo.db)
}
