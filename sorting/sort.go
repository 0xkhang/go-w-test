package sorting

// move the minimum number on top each iteration
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
