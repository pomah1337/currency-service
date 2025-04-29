package db

import (
	"currencyService/currency/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

func InitConnection(cfg config.DbCfg) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.Url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(10 * time.Minute)

	return db, nil
}
