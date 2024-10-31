package database

import (
	"fmt"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/qytela/example-project-layout/internal/pkg/logger"
)

var (
	db   *sqlx.DB
	once sync.Once
)

func InitializeDB() *sqlx.DB {
	once.Do(func() {
		var (
			DB_HOST = os.Getenv("DB_HOST")
			DB_PORT = os.Getenv("DB_PORT")
			DB_USER = os.Getenv("DB_USER")
			DB_PASS = os.Getenv("DB_PASS")
			DB_NAME = os.Getenv("DB_NAME")
		)

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

		conn, err := sqlx.Connect("postgres", dsn)
		if err != nil {
			logger.MakeLogEntry(nil).Panic("failed to connect database: ", err)
		}

		db = conn

		logger.MakeLogEntry(nil).Info("database connected successfully")
	})

	return db
}
