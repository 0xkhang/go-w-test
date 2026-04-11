package sorting

func BubbleSort(nums []int) []int {
	n := len(nums)

	for k := range n {
		swap := false
		for i := 0; i < n-k-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = Swap(nums[i], nums[i+1])
				swap = true
			}
		}
		if !swap {
			return nums
		}
	}
	return nums
}
