package sequence

func Sum(nums []int) int { 
    sum := 0
    for _, num := range nums {
       sum += num 
    }
    return sum
}

func SumAll(nums ...[]int) []int {
    result := make([]int, 0)
    for _, num := range nums {
        sum := 0    
        for _, v := range num {
            sum += v
        }
        result = append(result, sum)
    }
    return result
}
