package recursion

import (
	"testing"
)

func TestFibRecursive(t *testing.T) {
	t.Run("normal num", func(t *testing.T) {
		got := Fib(12)
		want := 144

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("larger num", func(t *testing.T) {
		got := Fib(40)
		want := 102334155

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("negative value", func(t *testing.T) {
		got := Fib(-1)
		want := 0

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("with memoi", func(t *testing.T) {
        mp := make(map[int]int, 0)
		got := FibMemoi(12, mp)
		want := 144

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestFibIterative(t *testing.T) {
	t.Run("test fib seq iteratively", func(t *testing.T) {
		got := FibIterative(12)
		want := 144

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test fib seq iteratively alternative approach", func(t *testing.T) {
		got := FibIterative(40)
		want := 102334155

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("negative value - Iterative", func(t *testing.T) {
		got := FibIterative(-1)
		want := 0

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
