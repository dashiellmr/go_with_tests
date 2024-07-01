package arrays

func Summation(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, array := range numbersToSum {
		sums = append(sums, Summation(array))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, array := range numbersToSum {
		if len(array) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := array[1:]
		sums = append(sums, Summation(tail))
	}
	return sums
}
