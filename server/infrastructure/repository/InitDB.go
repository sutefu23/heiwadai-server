package repository

import (
	"database/sql"
	"fmt"
	"os"
	"server/infrastructure/logger"

	_ "github.com/lib/pq"
)

var Conn *sql.DB

func InitDB() *sql.DB {
	var err error

	if Conn != nil {
		return Conn
	}

	// err = godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("failed to load .env file: ", err)
	// 	return nil, err
	// }

	user := os.Getenv("PSQL_USER")
	password := os.Getenv("PSQL_PASS")
	host := os.Getenv("PSQL_HOST")
	port := os.Getenv("PSQL_PORT")
	database := os.Getenv("PSQL_DBNAME")

	Conn, err = sql.Open("postgres",
		fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database))

	if err != nil {
		logger.Fatalf("OpenError: %v", err)
		panic("DB couldn't initialize!")
	}

	if err = Conn.Ping(); err != nil {
		logger.Fatalf("PingError: %v", err)
		panic("DB couldn't initialize!")
	}

	return Conn
}