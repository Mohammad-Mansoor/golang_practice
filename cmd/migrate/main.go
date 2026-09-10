package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	"github.com/Mohammad-Mansoor/go-api/internal/config"
)

func main() {
	fmt.Println(os.Args)
	cfg := config.MustLoad()
	if len(os.Args) < 2 {
		panic("usage: migrate <up | down>")
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd: %v", err)
	}
	migrationsDir := filepath.Join(wd, "migrations")
	log.Printf("migrations dir: %s", migrationsDir)

	d, err := iofs.New(os.DirFS(migrationsDir), ".")
	if err != nil {
		log.Fatalf("iofs.New: %v", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, cfg.DB_URL)
	if err != nil {
		log.Fatalf("migration.new: %v", err)
	}

	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err := m.Steps(-1); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Unknow command: %s", os.Args[1])
	}
	log.Println("Running Migration")
}
