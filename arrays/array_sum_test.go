package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("got %d expected %d given, %v", got, expected, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{3, 4, 5, 6, 7, 5}

		got := Sum(numbers)
		expected := 30

		if got != expected {
			t.Errorf("got %d expected %d given, %v", got, expected, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 5})
	expected := []int{3, 8}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %d expected %d ", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d ", got, want)
		}
	}

	t.Run("sum of elements expect first one", func(t *testing.T) {
		got := SumAllTails([]int{3, 4}, []int{3, 9})
		want := []int{4, 9}

		checkSums(t, got, want)
	})

	t.Run("empty slice returns 0", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 5})
		want := []int{0, 5}

		checkSums(t, got, want)
	})

}
