package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMsg := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Sam")
		want := "Hello, Sam"

		assertCorrectMsg(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is passed", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMsg(t, got, want)
	})

}
