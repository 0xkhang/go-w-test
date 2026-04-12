package recursion

import (
	"testing"
)

type Person struct {
    age int
    name string
}

func TestFib(t *testing.T) {
	t.Run("normal num", func(t *testing.T) {
		got := Fib(12)
		want := 144

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("normal num", func(t *testing.T) {
		got := Fib(0)
		want := 1

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
