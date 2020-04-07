package hello

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Chris", "en"), "hello Chris")
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		assertCorrectMessage(t, Hello("", "en"), "hello world")
	})

	t.Run("in German", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Fritz", "de"), "Hallo Fritz")
	})

	t.Run("in Spanish", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Torres", "es"), "Ola Torres")
	})
}
