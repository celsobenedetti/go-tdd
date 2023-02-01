package arrays

import (
	"testing"
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
