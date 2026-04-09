package sequence

import (
	"slices"
	"testing"
)

func TestFlip(t *testing.T) {
	t.Run("Flip slice", func(t *testing.T) {
		got := Flip([]int{1, 2, 3})
		want := []int{3, 2, 1}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestFlipRange(t *testing.T) {
	t.Run("Flip from range start to end", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 6, 7}
		got := FlipRange(nums, 0, 3)
		want := []int{4, 3, 2, 1, 6, 7}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
