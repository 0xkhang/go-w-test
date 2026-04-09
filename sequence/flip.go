package sequence

// Flip (reverse) the whole slice of integers
func Flip(nums []int) []int {
    s := NewStack()           
    result := make([]int, 0)

    for _, num := range nums {
        s.Push(num);
    }

    for !s.IsEmpty() {
        result = append(result, s.Top())
        s.Pop()
    }
    return result
}

// takes in an slice of integers, flip range start to end and returns the modified array 
func FlipRange(nums []int, start, end int) []int {
    result := make([]int, 0)     

    for i := start; i <= end; i++ {
        result = append(result, nums[i] )
    }

    result = Flip(result)

    for i := end + 1; i < len(nums); i++ {
        result = append(result, nums[i])
    }
    return result
}
