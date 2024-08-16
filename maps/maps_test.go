package main

import "testing"

// TestSearch tests the Search method of the Dictionary type.
func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		// Test for a word that exists in the dictionary.
		got, err := dictionary.Search("test")
		want := "this is just a test"

		// If an error is returned, fail the test.
		if err != nil {
			t.Fatal("should find added word:", err)
		}

		// Check if the returned definition matches the expected value.
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		// Test for a word that does not exist in the dictionary.
		_, err := dictionary.Search("unknown")
		if err == nil {
			t.Fatal("expected to get an error")
		}

		// Check if the error matches the expected error.
		assertError(t, err, ErrNotFound)
	})
}

// TestAdd tests the Add method of the Dictionary type.
func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		// Test adding a new word to the dictionary.
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		// Attempt to add the word to the dictionary.
		err := dictionary.Add(word, definition)

		// Ensure no error occurred.
		assertError(t, err, nil)
		// Verify the word was added correctly.
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		// Test adding a word that already exists in the dictionary.
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		// Attempt to add the word again.
		err := dictionary.Add(word, "new test")

		// Ensure the correct error is returned.
		assertError(t, err, ErrWordExists)
		// Verify the original definition was not overwritten.
		assertDefinition(t, dictionary, word, definition)
	})
}

// TestUpdate tests the Update method of the Dictionary type.
func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		// Test updating the definition of an existing word.
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		// Attempt to update the word's definition.
		err := dictionary.Update(word, newDefinition)

		// Ensure no error occurred.
		assertError(t, err, nil)
		// Verify the definition was updated correctly.
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		// Test updating a word that does not exist in the dictionary.
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		// Attempt to update the word, which does not exist.
		err := dictionary.Update(word, definition)

		// Ensure the correct error is returned.
		assertError(t, err, ErrWordDoesNotExist)
	})
}

// TestDelete tests the Delete method of the Dictionary type.
func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	// Delete the word from the dictionary.
	dictionary.Delete(word)

	// Attempt to search for the deleted word and ensure it returns an error.
	_, err := dictionary.Search(word)
	assertError(t, err, ErrNotFound)
}

// assertStrings is a helper function to compare two strings in tests.
// If the strings do not match, it reports an error.
func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// assertError is a helper function to compare two errors in tests.
// If the errors do not match, it reports an error.
func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected to get an error")
	}

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

// assertDefinition is a helper function to check the definition of a word in the dictionary.
// It searches for the word and compares the found definition to the expected definition.
func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
}
