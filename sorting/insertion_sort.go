package sorting

func InsertionSort(nums []int) []int {
	result := make([]int, 0)

	for _, num := range nums {
		result = SmartInsert(result, num)
	}

	return result
}

func SmartInsert(nums []int, data int) []int {
    n := len(nums)

    if n == 0 {
        nums = append(nums, data)
        return nums
    } 

    new_nums := make([]int, n+1)
    
    copy(new_nums, nums)

    pos := 0

    for pos < n && nums[pos] < data {
        pos++ 
    }

    for i := len(new_nums) - 2; i >= pos; i-- {
        new_nums[i+1] = new_nums[i]
    }
    new_nums[pos] = data
    return new_nums
}

func OptimalInsertionSort(nums []int) []int {
    n := len(nums)
    for i := 1; i < n; i++ {
        val := nums[i]    
        
        j := i - 1
        for j >= 0 && nums[j] > val {
            nums[j+1] = nums[j]
            j--;
        }
        nums[j+1] = val
    }
    return nums
}
