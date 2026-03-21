package main

import (
	"flag"
	"fmt"
	"log"
	"scaffoldy/pkg/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Tambahkan flag untuk force version jika database dirty
	forceVersion := flag.Int("force", -1, "Force a specific migration version (to fix dirty state)")
	flag.Parse()

	cfg := config.Load()

	dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	m, err := migrate.New(
		"file://./migrations",
		dsn,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Jika flag -force dipanggil, jalankan m.Force() lalu berhenti
	if *forceVersion != -1 {
		fmt.Printf("Forcing version to %d...\n", *forceVersion)
		if err := m.Force(*forceVersion); err != nil {
			log.Fatalf("Failed to force version: %v", err)
		}
		log.Println("Force success! Sekarang Anda bisa menjalankan migrasi secara normal kembali.")
		return
	}

	// Jalankan migrasi normal
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	log.Println("migration success")
}
