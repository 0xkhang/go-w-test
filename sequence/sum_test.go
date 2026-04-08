package sequence

import (
	"fmt"
	"reflect"
	"slices"
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

func TestSumAll(t *testing.T) {
    t.Run("sum all elements", func(t *testing.T) {
        got := SumAll([]int{1, 2}, []int{9, 8})
        want := []int{3, 17}

        if !slices.Equal(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    })

    t.Run("sum all empty slice", func(t *testing.T) {
        got := SumAll([]int{}, []int{-1, 8})
        want := []int{0, 7}

        if !slices.Equal(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    })
}

func TestSumAllTails(t *testing.T) {
    t.Run("test sum all tails", func(t *testing.T) {
        got := SumAllTails([]int{1,3,7}, []int{9,9,5,2})
        want := []int{10, 16}

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    })


    t.Run("test sum all tails have empty slice", func(t *testing.T) {
        got := SumAllTails([]int{1,3,7}, []int{9,9,5,2}, []int{})
        want := []int{10, 16, 0}

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    })
}

func ExampleSum() {
	sum := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(sum)
	// Output: 15
}
