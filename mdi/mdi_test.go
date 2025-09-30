package main

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "World")

	got := buffer.String()
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
