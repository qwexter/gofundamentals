package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	d := Dictionary{"test": "just a test"}

	t.Run("When given known key", func(t *testing.T) {
		got, _ := d.Search("test")
		want := "just a test"

		assert(t, got, want)
	})

	t.Run("when given unknown key", func(t *testing.T) {
		got, err := d.Search("unknown")
		want := ""

		assertError(t, err)
		assert(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("When key is unique", func(t *testing.T) {
		d := Dictionary{}
		err := d.Add("test", "just a test")

		assertNoError(t, err)

		want := "just a test"
		got, err := d.Search("test")

		assertNoError(t, err)
		assert(t, got, want)
	})
	t.Run("When try add existed key", func(t *testing.T) {
		d := Dictionary{}
		err := d.Add("test", "just a test")
		assertNoError(t, err)

		err = d.Add("test", "one more time")
		assertError(t, err)

		got, err := d.Search("test")

		want := "just a test"
		assertNoError(t, err)
		assert(t, got, want)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("When try to update existing key", func(t *testing.T) {
		word := "test"
		d := Dictionary{word: "just a test"}
		want := "a new defenition"

		err := d.Update(word, want)
		assertNoError(t, err)

		got, err := d.Search(word)

		assertNoError(t, err)
		assert(t, got, want)
	})

	t.Run("when there is no word for update", func(t *testing.T) {
		d := Dictionary{}
		err := d.Update("unknown", "should return error")
		assertError(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("when deleted word exists", func(t *testing.T) {
		word := "test"
		d := Dictionary{"test": "just a test"}

		err := d.Delete(word)
		assertNoError(t, err)

		_, err = d.Search(word)
		assertError(t, err)
	})

	t.Run("when no word for delete", func(t *testing.T) {
		word := "test"
		d := Dictionary{}

		err := d.Delete(word)
		assertError(t, err)
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("unexpected error", err)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Fatal("expected to get an error")
	}
}

func assert(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}
