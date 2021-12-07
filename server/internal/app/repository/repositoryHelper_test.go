package repository

import (
	"fmt"
	"strings"
	"testing"
)

func NewTestRepository(t *testing.T, dbURL string) (*Repository, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DbURL = dbURL

	r := NewRepository(config)
	if err := r.Open(); err != nil {
		t.Fatalf("Error while opening database: %s", err.Error())
	}

	return r, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := r.Db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE",
				strings.Join(tables, ", "))); err != nil {
				t.Fatalf("Error while executing db query: %s", err.Error())
			}
		}

		r.Close()
	}
}
