package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is the mapped value",
	}

	t.Run("search for known word", func(t *testing.T) {
		got, err := dictionary.Search("test")

		assertStrings(t, got, "this is the mapped value")

		if err != nil {
			t.Errorf("expected key to be found, but wasnt")
		}

	})

	t.Run("search for known word", func(t *testing.T) {
		_, err := dictionary.Search("nil")

		if err == nil {
			t.Errorf("expected to be key not found error")
		}

		assertStrings(t, err.Error(), ErrKeyNotFound.Error())
	})

}

func TestAdd(t *testing.T) {
	t.Run("add word to dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		word, definition := "mitocondria", "powerhouse of the cell"

		dictionary.Add(word, definition)
		got, err := dictionary.Search(word)

		if err != nil {
			t.Errorf("expected new word to be found in dictionary")
		}

		assertStrings(t, got, definition)
	})

	t.Run("try to add existing word", func(t *testing.T) {
		dictionary := Dictionary{"mitocondria": "powerhouse of the cell"}
		word, definition := "mitocondria", "powerhouse of the cell"

		err := dictionary.Add(word, definition)

		if err == nil {
			t.Errorf("expected word already exists error")
		}

		assertStrings(t, err.Error(), ErrWordAlreadyExists.Error())
	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
