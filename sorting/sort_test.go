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

    t.Run("Bubble Sort", func(t *testing.T) {
        nums := []int{8,4,3,6,5,2}
        got := BubbleSort(nums)
        want := []int{2,3,4,5,6,8}

        if !slices.Equal(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    })

    t.Run("Test Merge Sort 2 Sorted Slices", func(t *testing.T) {
        nums1 := []int{1,3,5,7}
        nums2 := []int{2,4,6,8}
        got := MergeSorted(nums1, nums2)
        want := []int{1,2,3,4,5,6,7,8}

        if !slices.Equal(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    })
}
