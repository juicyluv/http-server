package service

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ellywynn/http-server/internal/app/repository"
)

func NewTestService(t *testing.T, dbURL string) (*Service, func(...string)) {
	t.Helper()

	config := repository.NewConfig()
	config.DbURL = dbURL

	r := repository.NewRepository(config)
	if err := r.Open(); err != nil {
		t.Fatalf("Error while opening database: %s", err.Error())
	}

	s := NewService(r)

	s.User = NewUserService(&r.User)
	s.Auth = NewAuthService(&r.Auth)

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := r.Db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE",
				strings.Join(tables, ", "))); err != nil {
				t.Fatalf("Error while executing db query: %s", err.Error())
			}
		}

		r.Close()
	}
}
