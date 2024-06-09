package container

import (
	"context"
	"errors"
	"fmt"
	"log"
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
		log.Fatal("failed to setup container", err)
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

	time.Sleep(time.Second)

	dbAddr := fmt.Sprintf("localhost:%s", p.Port())
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", DBUser, DBPass, dbAddr, DBName))
	if err != nil {
		return container, db, dbAddr, fmt.Errorf("failed to establish database connection: %v", err)
	}

	return container, db, dbAddr, nil
}

func (tdb *TestDatabase) RunMigrations() error {
	migrationsFolderLocation := "../../db/migrations"
	pathToMigrate := fmt.Sprintf("file://%s", migrationsFolderLocation)
	connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", DBUser, DBPass, tdb.DBAddr, DBName)

	m, err := migrate.New(pathToMigrate, connectionString)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

func (tdb *TestDatabase) RollbackMigrations() error {
	migrationsFolderLocation := "../../db/migrations"
	pathToMigrate := fmt.Sprintf("file://%s", migrationsFolderLocation)
	connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", DBUser, DBPass, tdb.DBAddr, DBName)

	m, err := migrate.New(pathToMigrate, connectionString)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Down(); err != nil {
		return err
	}

	return nil
}
