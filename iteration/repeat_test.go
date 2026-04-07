package iteration

import "testing"

func TestRepeat1(t *testing.T) {
	repeated := Repeat1("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeat2(t *testing.T) {
	repeated := Repeat2("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat1(b *testing.B) {
	for b.Loop() {
		Repeat1("a", 100)
	}
}

func BenchmarkRepeat2(b *testing.B) {
	for b.Loop() {
		Repeat2("a", 100)
	}
}
