package sorting

import (
	"slices"
	"testing"
)

func TestSort(t *testing.T) {
    t.Run("Selection Sort", func(t *testing.T) {
        nums := []int{8,4,3,6,5,2}
        got := SelectionSort(nums)
        want := []int{2,3,4,5,6,8}

        if !slices.Equal(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    })

    t.Run("Reverse Selection Sort", func(t *testing.T) {
        nums := []int{8,4,3,6,5,2}
        got := ReverseSelectionSort(nums)
        want := []int{2,3,4,5,6,8}

        if !slices.Equal(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    })

    // t.Run("Insertion sort", func(t *testing.T) {
    //     nums := []int{8,4,3,6,5,2}
    //     got := InsertionSort(nums)
    //     want := []int{2,3,4,5,6,8}
    //
    //     if !slices.Equal(got, want) {
    //         t.Errorf("got %v want %v", got, want)
    //     }
    // })
}
