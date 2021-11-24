package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_HandleHello(t *testing.T) {
	s := NewServer(NewDefaultConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, "hello", rec.Body.String())
}
