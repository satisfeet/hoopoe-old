package httpd

import (
	"errors"
	"net/http"
	"testing"

	"github.com/satisfeet/hoopoe/store"
)

var (
	ErrTest = errors.New("a test error")
)

func TestErrorCode(t *testing.T) {
	if v := ErrorCode(ErrTest); v != http.StatusInternalServerError {
		t.Errorf("Expected to return 500 but had %d.\n", v)
	}
	if v := ErrorCode(store.ErrInvalidQuery); v != http.StatusBadRequest {
		t.Errorf("Expected to return 400 but had %d.\n", v)
	}
}
