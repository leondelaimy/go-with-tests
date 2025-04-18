package main

import "testing"

func TestHello(t *testing.T){
	
	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Leon", "")
		want := "Hello, Leon!"
		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})

	t.Run("say hello world when supplied empty string", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in spanish", func(t *testing.T){
		got := Hello("Leon", "Spanish")
		want := "Hola, Leon!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in french", func(t *testing.T){
		got := Hello("Leon", "French")
		want := "Bonjour, Leon!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}