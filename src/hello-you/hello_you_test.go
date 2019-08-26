package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := Hello("Nameless")
	want := "Hello Nameless"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
