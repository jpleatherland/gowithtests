package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people without language specified", func(t *testing.T) {
		got := Hello("", "Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string and language is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in spanish", func(t *testing.T){
		got := Hello("spanish", "Chris")
		want := "Hola, Chris"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
