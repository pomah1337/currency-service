package main

import (
	"currencyService/currency/internal/migrations"
	"flag"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	dbURL = "postgres://user:password@localhost:5432/currency?sslmode=disable"
)

func main() {
	action := flag.String("action", "up", "Migration action: (up, down, version)")
	version := flag.Uint("version", 0, "Version of database")
	path := flag.String("path", "currency/internal/migrations", "Path to migration files")
	flag.Parse()

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalf("connect db: %v", err)
	}
	defer db.Close()

	m, err := migrations.NewMigrator(db, *path)
	if err != nil {
		log.Fatalf("create migration: %v", err)
	}

	switch *action {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "version":
		err = m.SetVersion(*version)
	default:
		log.Fatalf("unknown action: %s", *action)
	}

	if err != nil {
		log.Fatalf("migration %s failed: %v", *action, err)
	}
}
