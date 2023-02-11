package generics_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/celso-patiri/go-tdd/generics"
)

func TestStack(t *testing.T) {

	t.Run("test 100 pushes and pops with ints", func(t *testing.T) {
		stack := generics.Stack[int]{}

		for i := 0; i < 100; i++ {
			stack.Push(i)
		}

		for i := 99; i >= 0; i-- {
			got, _ := stack.Pop()

			if got != i {
				t.Errorf("got %d wanted %d", got, i)
			}
		}
	})

	t.Run("test 100 pushes and pops with strings", func(t *testing.T) {
		stack := generics.Stack[string]{}

		for i := 0; i < 100; i++ {
			stack.Push(fmt.Sprint(i))
		}

		for i := 99; i >= 0; i-- {
			got, _ := stack.Pop()

			if got != fmt.Sprint(i) {
				t.Errorf("got %s wanted %d", got, i)
			}
		}
	})

	t.Run("test 100 concurrent pushes", func(t *testing.T) {
		stack := generics.Stack[int]{}

		wg := sync.WaitGroup{}
		wg.Add(100)

		for i := 0; i < 100; i++ {
			go func(value int) {
				stack.Push(value)
				wg.Done()
			}(i)
		}

		wg.Wait()

		if stack.Len() != 100 {
			t.Errorf("got %d wanted %d", stack.Len(), 100)
		}
	})

	t.Run("test 100 concurrent pops", func(t *testing.T) {
		stack := generics.Stack[int]{}

		for i := 0; i < 100; i++ {
			stack.Push(i)
		}

		wg := sync.WaitGroup{}
		wg.Add(100)

        for i := 0; i < 100; i++ {
            go func() {
                got, _ := stack.Pop()
                fmt.Println(got)
                wg.Done()
            }()
        }

		wg.Wait()
	})

}

func BenchmarkStack(b *testing.B) {
	stack := generics.Stack[int]{}
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}

	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}
