package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 3)
	expected := 5

	if sum != expected {
		t.Errorf("expected %d, got %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(3, 4)
	fmt.Println(sum)
	//Output: 7
}
