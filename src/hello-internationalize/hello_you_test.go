package main

import "testing"

func TestHelloWorld(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people in english", func(t *testing.T) {
		got := Hello("Nameless", "English")
		want := "Hello Nameless"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people in spanish", func(t *testing.T) {
		got := Hello("Nameless", "Spanish")
		want := "Hola Nameless"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people in french", func(t *testing.T) {
		got := Hello("Nameless", "French")
		want := "Bonjour Nameless"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say default salute", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
}
