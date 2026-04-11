package sorting

/*
The idea of the algorithm is to find the smallest unsorted element and add it to the end of the sorted list
*/
func SelectionSort(nums []int) []int {
	n := len(nums)
	for i := range n {
		min := i
		for j := i; j < n; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[min], nums[i] = nums[i], nums[min]
	}
	return nums
}

/*
Instead of moving the smallest unsorted element, we move the largest
*/
func ReverseSelectionSort(nums []int) []int {
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		max := i
		for j := i; j >= 0; j-- {
			if nums[j] > nums[max] {
				max = j
			}
		}
		nums[max], nums[i] = nums[i], nums[max]
	}
	return nums
}
