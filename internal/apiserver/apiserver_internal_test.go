package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleHello(t *testing.T) {
	s := NewAPIServer(NewConfig())
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)

	assert.Equal(t, rec.Body.String(), "Hello!")
}
