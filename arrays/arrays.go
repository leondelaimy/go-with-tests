package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	} 
	return sum
}

func SumAll(slices ...[]int) []int {
	numberOfSlices := len(slices)
	sums := make([]int, numberOfSlices)

	for i, numbers := range slices {
		sums[i] = Sum(numbers)
	}
	return sums
}

func SumAllTails(slices ...[]int) []int {
	var sums []int
	for _, numbers := range slices {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}	
	}
	return sums
}