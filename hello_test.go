package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	got := Hello("World")
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloYou(t *testing.T) {
	got := Hello("Celso")
	want := "Hello, Celso!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

