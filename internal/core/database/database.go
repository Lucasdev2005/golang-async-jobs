package database

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var Connection *pgxpool.Pool

func Connect() error {
	var err error
	Connection, err = pgxpool.Connect(context.Background(), getUrl())

	return err
}

func Close() {
	if Connection == nil {
		return
	}
	Connection.Close()
}

func getUrl() string {
	_, b, _, _ := runtime.Caller(0)
	var ProjectRootPath = filepath.Join(filepath.Dir(b), "../../../")
	godotenv.Load(ProjectRootPath + "/.env")

	var (
		dbPort     = os.Getenv("DB_PORT")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbUser     = os.Getenv("DB_USER")
		dbName     = os.Getenv("DB_NAME")
		dbHost     = os.Getenv("DB_HOST")
	)

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)
}
