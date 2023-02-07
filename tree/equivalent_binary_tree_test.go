package tree

import (
	"testing"

	"golang.org/x/tour/tree"
)

func TestSame(t *testing.T) {
	t1, t2 := tree.New(1), tree.New(1)

	got := Same(t1, t2)
	want := true

	if got != want {
		t.Errorf("Same: got %v expected %v", got, want)
	}
}

func TestWalk(t *testing.T) {
	tree := tree.New(1)
	ch := make(chan int)

	go Walk(tree, ch)

	for i := 1; i <= 10; i++ {
        value := <-ch
		if i != value {
			t.Errorf("Walk: got %d, expected %d", value, i)
		}
	}
}
