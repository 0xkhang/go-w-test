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

func BubbleSort(nums []int) []int {
	n := len(nums)

	for k := range n {
        swap := false
		for i := 0; i < n-k-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = Swap(nums[i], nums[i+1])
                swap = true;
			}
		}
        if !swap {
            return nums
        } 
	}
    return nums
}

func MergeSorted(nums1, nums2 []int) []int {
    result := make([]int, 0) 
    i := 0
    j := 0

    for i < len(nums1) && j < len(nums2) {
        if nums1[i] <= nums2[j] {
            result = append(result, nums1[i])
            i++
        } else {
            result = append(result, nums2[j])
            j++
        }
    }

    for i < len(nums1) {
        result = append(result, nums1[i])
        i++
    }

    for j < len(nums1) {
        result = append(result, nums2[j])
        j++
    }

    return result 
}

func Swap(a, b int) (int, int) {
	temp := a
	a = b
	b = temp
	return a, b
}
