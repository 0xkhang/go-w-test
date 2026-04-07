package sequence

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum with all positive integers", func(t *testing.T) {
		nums := []int{10, 20, 30, 40, 50}

		got := Sum(nums)
		want := 150

		if got != want {
			t.Errorf("got %d want %d given %v\n", got, want, nums)
		}
	})
    
	t.Run("sum with all negative integers", func(t *testing.T) {
		nums := []int{-10, -20, -30, -40, -50}

		got := Sum(nums)
		want := -150

		if got != want {
			t.Errorf("got %d want %d given %v\n", got, want, nums)
		}
	})
}

func ExampleSum() {
	sum := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(sum)
	// Output: 15
}
