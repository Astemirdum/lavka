package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Astemirdum/lavka/internal/config"
	"github.com/Astemirdum/lavka/migrations"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
)

func NewPostgresDB(cfg *config.DB) (*sqlx.DB, error) {
	//NOTE: migrate-job
	// if err := migrateSchema(cfg); err != nil {
	//	return nil, err
	//}
	return newDB(cfg)
}

func newDB(cfg *config.DB) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", newDSN(cfg))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("db master ping: %w", err)
	}
	return db, nil
}

func newDSN(cfg *config.DB) string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.NameDB, cfg.Password)
}

//nolint:unused
func migrateSchema(cfg *config.DB) error {
	dsn := newDSN(cfg)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		return fmt.Errorf("migrateSchema ping: %w", err)
	}

	goose.SetBaseFS(migrations.MigrationFiles)
	if err = goose.Up(db, "."); err != nil {
		return errors.Wrap(err, "goose run()")
	}
	return nil
}
