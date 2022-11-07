package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSsh(t *testing.T) {
	// assertCorrectMsg := func(t testing.TB, got, want string) {
	// 	t.Helper()
	// 	if got != want {
	// 		t.Errorf("got %q want %q", got, want)
	// 	}
	// }

	t.Run("generate public key in openssh format", func(t *testing.T) {
		got, err := generateKeyPair()
		if err != nil {
			t.Error(err)
		}
		want := "ssh-rsa"
		fmt.Println(got)

		if strings.Contains(got.publicKey, want) != true {
			t.Errorf("got %q want %q", *got, want)
		}
	})

}
