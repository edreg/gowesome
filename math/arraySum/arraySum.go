package arraySum

func ArraySum(numbers []int) int {

	result := 0

	//length := len(numbers)
	//for i := 0; i < length; i++ {
	//	result += numbers[i]
	//}

	// _ aka Mr. Blanky
	for _, number := range numbers {
		result += number
	}

	return result
}

func SumNArrays(intArray ...[]int) []int {
	//result := make([]int, len(intArray))
	//
	//for i, numbers := range intArray {
	//	result[i] = ArraySum(numbers)
	//}
	//
	//return result

	var sums []int
	for _, numbers := range intArray {
		sums = append(sums, ArraySum(numbers))
	}

	return sums
}

func SumArrayTails(intArray ...[]int) []int {
	var sums []int
	for _, numbers := range intArray {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			//sums = append(sums, ArraySum(numbers[1:len(numbers)]))
			sums = append(sums, ArraySum(numbers[1:]))
		}

	}

	return sums
}

func SumArrayValuesAtPosition(intArray ...[]int) []int {
	var sums []int
	for _, numbers := range intArray {

		for i, number := range numbers {
			if i > len(sums)-1 {
				sums = append(sums, 0)
			}
			sums[i] += number
		}
	}

	return sums
}
