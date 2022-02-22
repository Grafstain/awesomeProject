package array

func ArraySum(numbers []int) int {
	sum := 0
	//for i:=0;i <len(numbers);i++{
	//	sum += numbers[i]
	//}
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, ArraySum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, ArraySum(tail))
		}
	}

	return sums
}
