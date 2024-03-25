package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Zé", "English")
		want := "Hello, Zé!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Pedro", "Portuguese")
		want := "Olá, Pedro!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Pedro", "French")
		want := "Bonjour, Pedro!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// makes it so the line number on failed test appears on original function instead of helper
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
