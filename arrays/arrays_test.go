package arrays

import (
	"testing"

	"github.com/matryer/is"
)

func TestSum(t *testing.T) {
	t.Run("sum of 5 length array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		sum := Sum(numbers)
		expected := 15

		if sum != expected {
			t.Errorf("expected sum of %v to be %d, but got %d", numbers, expected, sum)
		}
	})

	t.Run("sum of slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		sum := Sum(numbers)
		expected := 15

		if sum != expected {
			t.Errorf("expected sum of %v to be %d, but got %d", numbers, expected, sum)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum 2 different slices", func(t *testing.T) {
		slices := [][]int{
			{1, 2, 3},
			{4, 5, 6},
		}

		sum := SumAll(slices)
		expected := 21

		if sum != expected {
			t.Errorf("expected sum of %v to be %d, but got %d", slices, expected, sum)
		}
	})
}

func TestReduce(t *testing.T) {
	is := is.New(t)

	t.Run("reduce a slice of int", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Reduce(numbers, func(a, b int) int {
			return a + b
		}, 0)

		want := 10
		is.Equal(got, want)
	})

}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		is := is.New(t)
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		is.True(found != -1)
		is.Equal(firstEvenNumber, 2)
	})
}
