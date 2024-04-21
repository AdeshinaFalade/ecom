package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/adeshinafalade/ecom/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	dbConfig := config.Envs
	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", dbConfig.DBUser, dbConfig.DBPassword, dbConfig.Host, dbConfig.Port, dbConfig.DBName)

	db, _ := sql.Open("postgres", connStr)
	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres", driver)
	// m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run

	// m, err := migrate.New(
	// 	"file://cmd/migrate/migrations",
	// 	connStr)

	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}

// to run migration
// migrate create -ext sql -dir cmd/migrate/migrations -seq create_order_items_table
//
//migrate -database 'postgres://postgres:shred@localhost:1999/ecom?sslmode=disable' -path cmd/migrate/migrations up

//export POSTGRESQL_URL='postgres://postgres:shred@localhost:1999/ecom?sslmode=disable&search_path=public'

//migrate -database 'postgres://postgres:shred@localhost:1999/ecom?sslmode=disable&search_path=public' -path cmd/migrate/migrations up

//migrate -database postgres://postgres:shred@localhost:1999/ecom?sslmode=disable -path cmd/migrate/migrations up

//migrate -database postgres://postgres:shred@localhost:1999/ecom?sslmode=disable -path cmd/migrate/migrations up
