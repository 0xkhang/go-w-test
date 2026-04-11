package sorting

import (
	"slices"
	"testing"
)

func TestSort(t *testing.T) {
	t.Run("Selection Sort", func(t *testing.T) {
		nums := []int{8, 4, 3, 6, 5, 2}
		got := SelectionSort(nums)
		want := []int{2, 3, 4, 5, 6, 8}

		Compare(t, got, want)
	})

	t.Run("Reverse Selection Sort", func(t *testing.T) {
		nums := []int{8, 4, 3, 6, 5, 2}
		got := ReverseSelectionSort(nums)
		want := []int{2, 3, 4, 5, 6, 8}

		Compare(t, got, want)
	})

	t.Run("Bubble Sort", func(t *testing.T) {
		nums := []int{8, 4, 3, 6, 5, 2}
		got := BubbleSort(nums)
		want := []int{2, 3, 4, 5, 6, 8}

		Compare(t, got, want)
	})

	t.Run("Test Merge Sort 2 Sorted Slices", func(t *testing.T) {
		nums1 := []int{1, 3, 5, 7}
		nums2 := []int{2, 4, 6, 8}
		got := MergeSorted(nums1, nums2)
		want := []int{1, 2, 3, 4, 5, 6, 7, 8}

		Compare(t, got, want)
	})

	t.Run("Test InsertionSort", func(t *testing.T) {
		nums := []int{5, 4, 2, 1}
		got := InsertionSort(nums)
		want := []int{1, 2, 4, 5}

		Compare(t, got, want)
	})

	t.Run("InsertionSort - Duplicate Numbers", func(t *testing.T) {
		nums := []int{3, 3, 1, 2, 3}
		got := InsertionSort(nums)
		want := []int{1, 2, 3, 3, 3}

		Compare(t, got, want)
	})

	t.Run("InsertionSort - Insert at End", func(t *testing.T) {
		nums2 := []int{5, 4, 3, 2, 1}
		got2 := InsertionSort(nums2)
		want2 := []int{1, 2, 3, 4, 5}

		Compare(t, got2, want2) // This should fail!
	})

	t.Run("InsertionSort - Single Element", func(t *testing.T) {
		nums := []int{42}
		got := InsertionSort(nums)
		want := []int{42}

		Compare(t, got, want)
	})

	t.Run("Test Optimal", func(t *testing.T) {
		nums := []int{5, 4, 2, 1}
		got := OptimalInsertionSort(nums)
		want := []int{1, 2, 4, 5}

		Compare(t, got, want)
	})


	t.Run("Test Optimal", func(t *testing.T) {
		nums := []int{5, 4, 2, 1, 3, 3, 9, 7}
		got := OptimalInsertionSort(nums)
		want := []int{1, 2, 3, 3 ,4, 5, 7, 9}

		Compare(t, got, want)
	})
}

func Compare(t *testing.T, got, want []int) {
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
