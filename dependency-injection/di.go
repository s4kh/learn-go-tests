package dependencyinjection

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, value string) {
	fmt.Fprintf(writer, "Hello, %s", value)
}
