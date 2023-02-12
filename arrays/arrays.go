package arrays

func Sum(numbers []int) (sum int) {
	add := func(a, b int) int { return a + b }
	return Reduce(numbers, add, 0)
}

// Sum calculates the total from a slice of numbers.
func SumAll(slices [][]int) (sum int) {
	for _, slice := range slices {
		sum += Sum(slice)
	}
	return
}

// SumAllTails calculates the sums of all but the first number given a collection of slices.
func SumAllTails(slices ...[]int) (results []int) {
	sumAll := func(acc, x []int) []int {
        if len(x) == 0 {
            return append(acc, 0)
        }else {
            return append(acc, Sum(x[1:]))
        }
	}
    return Reduce(slices, sumAll, []int{})
}

func Reduce[A, B any](collection []A, reducer func(B, A) B, initalValue B) B {
	acc := initalValue
	for _, item := range collection {
		acc = reducer(acc, item)
	}
	return acc
}

func Find[A any](collection []A, predicate func(A) bool) (item A, idx int)  {
    for i, item := range collection {
        if predicate(item) {
            return item, i
        }
    }
    return item, -1
    
}
