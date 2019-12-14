package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to peaple", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("", func(t *testing.T) {
		got := Hello("Elodie", spanish)
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})
	t.Run("", func(t *testing.T) {
		got := Hello("Elodie", french)
		want := "Bonjure, Elodie"

		assertCorrectMessage(t, got, want)
	})

}
