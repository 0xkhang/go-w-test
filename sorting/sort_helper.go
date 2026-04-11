package sorting

// helper for swap
func Swap(a, b int) (int, int) {
	temp := a
	a = b
	b = temp
	return a, b
}
