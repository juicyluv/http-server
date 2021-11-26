package service_test

import (
	"os"
	"testing"
)

var dbURL string

func TestMain(m *testing.M) {
	dbURL = os.Getenv("TEST_DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:qwerty@localhost:5437/travels_test?sslmode=disable"
	}

	os.Exit(m.Run())
}
