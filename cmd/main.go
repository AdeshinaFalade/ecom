package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/adeshinafalade/ecom/cmd/api"
	"github.com/adeshinafalade/ecom/config"
	"github.com/adeshinafalade/ecom/db"
)

func main() {
	dbConfig := config.Envs
	connStr := fmt.Sprintf("user=%v password=%v port=%v dbname=%v host=%v sslmode=disable", dbConfig.DBUser, dbConfig.DBPassword, dbConfig.Port, dbConfig.DBName, dbConfig.Host)
	db, err := db.NewMySQLStorage(connStr)
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)
	server := api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("DB connected successfully!")
}
