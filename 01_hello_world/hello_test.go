package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Advait", "")
		want := "Hello, Advait"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say \"Hello, World\" when empty string is given", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Ezio", "French")
		want := "Bonjour, Ezio"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Hindi", func(t *testing.T) {
		got := Hello("Tejas", "Hindi")
		want := "Namaste, Tejas"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
