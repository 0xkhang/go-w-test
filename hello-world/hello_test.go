package helloworld

import "testing"

func TestHello(t *testing.T) {
    // argument provided
	t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris", "english")
		want := "Hello, Chris"

        assertCorrectMessage(t, got, want)
	})

    // no argument provided
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "viet")
		want := "Chao, World"

        assertCorrectMessage(t, got, want)
	})

    // Spanish
	t.Run("spanish", func(t *testing.T) {
		got := Hello("", "ESPANOL")
		want := "Ola, World"

        assertCorrectMessage(t, got, want)
	})
    // no argument provided
	t.Run("Chinese", func(t *testing.T) {
		got := Hello("Van", "Chinese")
		want := "Cao ni ma, Van"

        assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
    if got != want {
       t.Errorf("got %q want %q", got, want)  
    }
}
