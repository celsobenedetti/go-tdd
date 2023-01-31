package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to Celso", func(t *testing.T) {
		got := Hello("Celso")
		want := "Hello, Celso!"

		assertCorrectMesssage(t, got, want)
	})

	t.Run("saying hello with empty string", func(t *testing.T) {
		got := Hello()
		want := "Hello, World!"

		assertCorrectMesssage(t, got, want)
	})

	t.Run("saying Hello to Jose in spanish", func(t *testing.T) {
		got := Hello("Jose", "sp")
		want := "Hola, Jose!"

		assertCorrectMesssage(t, got, want)
	})

	t.Run("saying Hello to Pierre in french", func(t *testing.T) {
		got := Hello("Pierre", "fr")
		want := "Bonjour, Pierre!"

		assertCorrectMesssage(t, got, want)
	})

}

func assertCorrectMesssage(t testing.TB, got, want string) {
	//used to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
