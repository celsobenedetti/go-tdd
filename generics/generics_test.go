package generics

import "testing"

func TestAssertFunction(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEquals(t, 1, 1)
		AssertNotEquals(t, 1, 2)
	})
}

func AssertEquals[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}

}

func AssertNotEquals[T comparable](t testing.TB, got, want T) {
    t.Helper()

    if got == want {
        t.Errorf("did not want %+v", got)
    }
}
