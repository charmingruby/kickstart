package container

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	DBName = "test_db"
	DBUser = "test_user"
	DBPass = "test_password"
)

type TestDatabase struct {
	DB        *sqlx.DB
	DBAddr    string
	container testcontainers.Container
}

func NewPostgresTestDatabase() *TestDatabase {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	container, db, dbAddr, err := createContainer(ctx)
	if err != nil {
		log.Fatal("failed to setup test", err)
	}

	if err := runMigrations(dbAddr); err != nil {
		log.Fatal("failed to perform db migration", err)
	}
	cancel()

	return &TestDatabase{
		container: container,
		DB:        db,
		DBAddr:    dbAddr,
	}
}

func (tdb *TestDatabase) Teardown() {
	tdb.DB.Close()
	_ = tdb.container.Terminate(context.Background())
}

func createContainer(ctx context.Context) (testcontainers.Container, *sqlx.DB, string, error) {
	var env = map[string]string{
		"POSTGRES_PASSWORD": DBPass,
		"POSTGRES_USER":     DBUser,
		"POSTGRES_DB":       DBName,
	}
	var port = "5432/tcp"

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:14-alpine",
			ExposedPorts: []string{port},
			Env:          env,
			WaitingFor:   wait.ForLog("database system is ready to accept connections"),
		},
		Started: true,
	}
	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return container, nil, "", fmt.Errorf("failed to start container: %v", err)
	}

	p, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return container, nil, "", fmt.Errorf("failed to get container external port: %v", err)
	}

	slog.Info("postgres container ready and running at port: " + p.Port())
	time.Sleep(time.Second)

	dbAddr := fmt.Sprintf("localhost:%s", p.Port())
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", DBUser, DBPass, dbAddr, DBName))
	if err != nil {
		return container, db, dbAddr, fmt.Errorf("failed to establish database connection: %v", err)
	}

	return container, db, dbAddr, nil
}

func runMigrations(dbAddr string) error {
	migrationsFolderLocation := "../../db/migrations"
	pathToMigrate := fmt.Sprintf("file://%s", migrationsFolderLocation)
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", DBUser, DBPass, dbAddr, DBName)

	println(pathToMigrate)

	m, err := migrate.New(pathToMigrate, databaseURL)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	slog.Info("migration done")

	return nil
}
