package repository

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

const (
	ROOT_DIR       = "../.."
	MYSQL_DATABASE = "sample_todo_app"
	MYSQL_USER     = "user"
	MYSQL_PASSWORD = "password"
)

var (
	MYSQL_HOST string
	MYSQL_PORT string
)

func TestMain(m *testing.M) {
	ctx := context.TODO()
	container, err := mysql.Run(ctx,
		"mysql:8.4",
		mysql.WithDatabase(MYSQL_DATABASE),
		mysql.WithUsername(MYSQL_USER),
		mysql.WithPassword(MYSQL_PASSWORD),
		mysql.WithScripts(filepath.Join(ROOT_DIR, "testdata/sql", "todos.sql")),
	)
	defer func() {
		if err := testcontainers.TerminateContainer(container); err != nil {
			log.Printf("failed to terminate container: %v", err)
		}
	}()
	if err != nil {
		log.Fatalf("failed to start container: %v", err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		log.Fatalf("failed to get containers host: %v", err)
	}
	MYSQL_HOST = host

	port, err := container.MappedPort(ctx, "3306")
	if err != nil {
		log.Fatalf("failed to get contaier's port: %v", err)
	}
	MYSQL_PORT = port.Port()

	statusCode := m.Run()
	os.Exit(statusCode)
}
