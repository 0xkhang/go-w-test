package recursion

import "testing"

func TestExp(t *testing.T) {
	t.Run("test exponential iteratively", func(t *testing.T) {
		got := ExpIterative(2, 3)
		want := 8

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test exponential recursively", func(t *testing.T) {
		got := ExpIterative(2, 4)
		want := 16

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test exponential recursively", func(t *testing.T) {
		got := ExpIterative(3, 3)
		want := 27

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test exponential smart for even", func(t *testing.T) {
		got := ExpSmart(2, 8)
		want := 256

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
