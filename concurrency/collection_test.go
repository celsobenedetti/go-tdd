package concurrency

import (
	"fmt"
	"sync"
	"testing"
)

func TestNewCollection(t *testing.T) {
	c := NewCollection()

	if c.Data == nil {
		t.Errorf("NewCollection: expected collection not to be nil")
	}
}

func TestAdd(t *testing.T) {

	t.Run("add simple key value pair", func(t *testing.T) {

		c := NewCollection()

		key, value := "key", "value"

		c.Add(key, value)

		got, ok := c.Data[key]

		if !ok || got != value {
			t.Errorf("Add: expected key, value pair to be added to map")
		}
	})

	t.Run("add with multiple parallel routines", func(t *testing.T) {
		c := NewCollection()

        count := 1000

        var wg sync.WaitGroup
		for i := 0; i < count; i++ {
            wg.Add(1)

			go func(key string) {
                defer wg.Done()
				c.Add(key, key)
			}(fmt.Sprint(i))
		}

        wg.Wait()

        for i := 0; i < count; i++ {
            if !c.Has(fmt.Sprint(i)){
                t.Errorf("Add: expected map to have key %q", i)
            }
        }
	})
}

func TestHas(t *testing.T) {
	c := NewCollection()

	key := "key"

	c.Add(key, "any")

	has := c.Has(key)

	if !has {
		t.Errorf("Expected %q to be present in map", key)
	}

	has = c.Has("banana")

	if has {
		t.Errorf("Expected banana not to be present in map")
	}
}
