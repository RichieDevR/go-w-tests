package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Richard", "English")

		want := "Hello, Richard"

		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to world", func(t *testing.T) {
		got := Hello("", "")

		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Ricardo", "Spanish")

		want := "Hola, Ricardo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Richard", "French")

		want := "Bonjour, Richard"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Ricardo", "Portuguese")

		want := "Olá, Ricardo"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
