package arrays

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(slices [][]int) (sum int) {
	for _, slice := range slices {
        sum += Sum(slice)
	}
    return
}
