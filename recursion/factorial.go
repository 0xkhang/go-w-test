package recursion

func Factorial(num int) int {
    if num == 0 || num == 1 {
        return num
    }

    return num * Factorial(num - 1)
}
