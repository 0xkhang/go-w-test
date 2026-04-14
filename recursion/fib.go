package recursion

// Iterative Approach
func FibIterative(n int) int {
	if n < 0 {
		return 0
	}

	if n <= 1 {
		return n
	}

	f1 := 0
	f2 := 1
	var f int
	for i := 2; i <= n; i++ {
		f = f1 + f2
		f1 = f2
		f2 = f
	}
	return f
}

// Recursive Approach
func Fib(n int) int {
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}

// Recursive Approach with Memoization (Top-Down)
func FibMemoi(n int, mp map[int]int) int {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	if val, ok := mp[n]; ok {
		return val
	}

	mp[n] = FibMemoi(n-1, mp) + FibMemoi(n-2, mp)
	return mp[n]
}
