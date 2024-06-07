package container

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	testDB := NewPostgresTestDatabase()
	defer testDB.Teardown()
	os.Exit(m.Run())
}
