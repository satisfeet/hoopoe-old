package store

import "testing"

func TestConnection(t *testing.T) {
	if err := Open("localhost/test"); err != nil {
		t.Errorf("Expected not to have error %s\n", err)
	}
	if mongo == nil {
		t.Error("Expected to have mongo not nil")
	}
	Close()
}
