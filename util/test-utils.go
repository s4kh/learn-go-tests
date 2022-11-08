// Package util contains all the test related utility functions
package util

import "testing"

func AssertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func AssertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Error("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
