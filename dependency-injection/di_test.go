package dependencyinjection

import (
	"bytes"
	"testing"

	"github.com/s4kh/learn-go-tests/util"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	util.AssertString(t, got, want)
}
