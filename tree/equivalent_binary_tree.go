package tree

import (
	"golang.org/x/tour/tree"
)

// Walk the tree t sending all values from tree to channel ch
func Walk(t *tree.Tree, ch chan int) {
    if t == nil {
        return
    }

    Walk(t.Left, ch)
    ch <- t.Value
    Walk(t.Right, ch)
}

// Returns wheter trees t1 and t2 have the same values
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

    for i := 0; i < 10; i++ {
        if <-ch1 != <-ch2 {
            return false
        }
    }
    return true
}
