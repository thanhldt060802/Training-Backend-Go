package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

var DB *bun.DB

func ConnectDB() *bun.DB {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		AppConfig.DBUser, AppConfig.DBPassword, AppConfig.DBHost, AppConfig.DBPort, AppConfig.DBName,
	)

	pgdb, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to open PostgreSQL connection:", err)
	}

	db := bun.NewDB(pgdb, pgdialect.New())
	DB = db

	if err := pgdb.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Connected to PostgreSQL with Bun ORM!")
	return DB
}
