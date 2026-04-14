package recursion

// Iterative approach
// x^n = x * x - n times
// T(n) = O(n)

func ExpIterative(x, n int) int {
    res := 1
    for range n {
        res *= x
    }
    return res     
}

// Recursive approach to compute x^n
// Assume each operation takes 1 unit of time:
//   - Comparison (n == 0)      : 1 unit
//   - Multiplication (x * ...) : 1 unit
//   - Subtraction (n - 1)      : 1 unit
// Total per recursive call (excluding the recursive call itself): ~3 units
//
// Recurrence relation:
//   T(n) = T(n-1) + 3
//
// Solving by unrolling:
//   T(n) = T(n-2) + 6
//   T(n) = T(n-3) + 9
//   T(n) = T(n-k) + 3k
//
// Base case: T(0) = 1 (just the comparison and return)
// Set n - k = 0  →  k = n
//
// Therefore:
//   T(n) = T(0) + 3n
//   T(n) = 1 + 3n
//
// Asymptotic complexity:
//   T(n) = O(n)   (linear time complexity)

func ExpRecursive(x, n int) int {
    if n == 0 {
        return 1
    }
    return x * ExpRecursive(x, n-1)
}

func ExpSmart(x, n int) int {
    if n == 0 {
       return 1 
    }

    if n % 2 == 0 && n != 0 {
        num := ExpSmart(x, n/2)
        return num * num
    }

    return x * ExpSmart(x, n - 1)
}
