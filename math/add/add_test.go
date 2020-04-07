package add

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(1, 1)
	expected := 1 + 1

	if got != expected {
		t.Errorf("got %d want %d", got, expected)
	}
}
