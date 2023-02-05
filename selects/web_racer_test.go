package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("test fast vs slow server", func(t *testing.T) {
		slowServer := makeDelayedServer(20)
		fastServer := makeDelayedServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl, fastUrl := slowServer.URL, fastServer.URL

		got, _ := Racer(slowUrl, fastUrl)
		want := fastUrl

		if got != want {
			t.Errorf("Racer: got %q, expected %q", got, want)
		}
	})

	t.Run("returns error if requests exceeds timeout ", func(t *testing.T) {
		server := makeDelayedServer(12 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, time.Millisecond)

		if err == nil {
			t.Errorf("Racer: expected error but got none")
		}
	})

}

func makeDelayedServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}
