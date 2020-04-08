package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Me")

	got := buffer.String()
	want := "Hello Me"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
