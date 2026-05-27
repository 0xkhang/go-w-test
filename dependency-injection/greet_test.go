package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
    buffer := bytes.Buffer{}
    Greet(&buffer, "Khang")

    got := buffer.String()
    want := "Hello, Khang!"

    if got != want {
        t.Errorf("got %q wnat %q", got, want)
    }
}
