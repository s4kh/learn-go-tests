package arrays

func Sum(numbers []int) int {
	var sum int

	for _, val := range numbers {
		sum += val
	}

	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, numSlice := range numbers {
		sums = append(sums, Sum(numSlice))
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, numSlice := range numbers {
		if len(numSlice) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := numSlice[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
