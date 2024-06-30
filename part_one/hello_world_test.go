package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("if a string was input", func(t *testing.T) {
		got := Hello("Dashiell", "English")
		want := "Hello Dashiell"
		assertCorrectMessage(t, got, want)
	})
	t.Run("if no string was input", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("if a string was input and the language is spanish", func(t *testing.T) {
		got := Hello("Dashiell", "Spanish")
		want := "Hola Dashiell"
		assertCorrectMessage(t, got, want)
	})
	t.Run("if no string was input and the language is spanish", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola World!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
