package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment the counter 3 times leaves it at 3", func(t *testing.T) {
		c := NewCounter()
		c.Inc()
		c.Inc()
		c.Inc()

		assertCounterValue(t, c, 3)
	})

	t.Run("increment counter with 1000 parallel routines", func(t *testing.T) {
		c := NewCounter()
		var wg sync.WaitGroup

		threadCount := 1000
		wg.Add(threadCount)

		for i := 0; i < threadCount; i++ {
			go func() {
				c.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		assertCounterValue(t, c, threadCount)
	})
}

func assertCounterValue(t testing.TB, c *Counter, want int) {
	t.Helper()

	if c.Value() != want {
		t.Errorf("Counter: expeceted %d, got %d", want, c.Value())
	}
}
