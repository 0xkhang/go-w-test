package sequence

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(nums ...[]int) []int {
	numOfSlice := len(nums)
	result := make([]int, numOfSlice)
	for i, num := range nums {
		result[i] = Sum(num)
	}
	return result
}

func SumAllTails(nums ...[]int) []int {
	numOfSlice := len(nums)
	result := make([]int, numOfSlice)
	for i, num := range nums {
		if len(num) == 0 {
			result[i] = 0
		} else {
			result[i] = Sum(num) - num[0]
		}
	}
	return result
}
