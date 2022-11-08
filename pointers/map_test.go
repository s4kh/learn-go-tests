package pointers

import (
	"testing"

	"github.com/s4kh/learn-go-tests/util"
)

func TestSearch(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		dict := Dictionary{"testing": "val", "testing2": "val2"}

		got, err := dict.Search("testing2")
		want := "val2"

		if err != nil {
			t.Errorf("expecting no error")
		}

		util.AssertString(t, got, want)
	})

	t.Run("non existing key", func(t *testing.T) {
		dict := Dictionary{"test": "woohoo"}

		got, err := dict.Search("test2")

		util.AssertString(t, got, "")
		util.AssertError(t, err, ErrNonExistingKey)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new key", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test3", "val3")

		want := "val3"
		got, err := dictionary.Search("test3")

		if err != nil {
			t.Fatal("should find the value:", err)
		}

		util.AssertString(t, got, want)
	})

	t.Run("existing key", func(t *testing.T) {
		dictionary := Dictionary{"test4": "val4"}
		err := dictionary.Add("test4", "val4")

		util.AssertError(t, err, ErrKeyExists)
	})

}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test5": "val5"}

	dictionary.Delete("test5")

	_, err := dictionary.Search("test5")

	util.AssertError(t, err, ErrNonExistingKey)
}
