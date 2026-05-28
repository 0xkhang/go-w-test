package mocking

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	expected_total_sleep := 3

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!
`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != expected_total_sleep {
		t.Errorf("not enough calls to sleeper. want %d got %d", expected_total_sleep, spySleeper.Calls)
	}
}
