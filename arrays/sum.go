package sum

// Sum numbers
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// lenghtOfNumbers := len(numbersToSum)
	// sums := make([]int, lenghtOfNumbers)
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// Sum independently each slice and returns a slice with each result
func SumAllTails(numbersToSum ...[]int) []int {
	var slices [][]int
	for _, slice := range numbersToSum {
		tail := slice[1:]
		slices = append(slices, tail)
	}
	// Spread operator to unroll the slice of slices
	return SumAll(slices...)
}
