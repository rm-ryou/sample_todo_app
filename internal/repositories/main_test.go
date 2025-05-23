package repositories

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/rm-ryou/sample_todo_app/internal/config"
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
	RoomRepo   *RoomRepository
	BoardRepo  *BoardRepository
	TodoRepo   *TodoRepository
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
		mysql.WithScripts(filepath.Join(ROOT_DIR, "testdata/sql", "schema.sql")),
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

	cfg := config.DB{
		Database: MYSQL_DATABASE,
		User:     MYSQL_USER,
		Password: MYSQL_PASSWORD,
		Host:     MYSQL_HOST,
		Port:     MYSQL_PORT,
	}
	db, err := SetupConnection(cfg)
	if err != nil {
		log.Fatalf("failed to connection db: %v", err)
	}
	defer db.Close()

	RoomRepo = NewRoomRepository(db)
	BoardRepo = NewBoardRepository(db)
	TodoRepo = NewTodoRepository(db)

	statusCode := m.Run()
	os.Exit(statusCode)
}
