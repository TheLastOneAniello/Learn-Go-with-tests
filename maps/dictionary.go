package main

// DictionaryErr is a custom error type that implements the error interface.
type DictionaryErr string

// Error method allows DictionaryErr to satisfy the error interface.
// It returns the error message as a string.
func (e DictionaryErr) Error() string {
	return string(e)
}

// Error constants for specific error scenarios.
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// Dictionary is a custom type that wraps around a map to provide additional methods.
type Dictionary map[string]string

// Search looks for a word in the dictionary.
// The function signature `Search(word string) (string, error)` means:
// - It takes a single input of type `string` (the word to search for).
// - It returns two values:
//  1. A `string` representing the definition of the word.
//  2. An `error` indicating whether something went wrong (e.g., if the word is not found).
//
// Returning an `error` is a common pattern in Go for indicating success or failure.
func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		// If the word is not found in the dictionary, return an empty string and an error.
		return "", ErrNotFound
	}
	// If the word is found, return the definition and a nil error.
	return value, nil
}

// Add inserts a word and its definition into the dictionary.
// The function signature `Add(word, definition string) error` means:
// - It takes two inputs of type `string` (the word and its definition).
// - It returns a single `error` value, which is nil if the operation is successful.
// Returning an `error` allows the function to signal if something went wrong,
// such as trying to add a word that already exists.
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		// Word does not exist, so add it to the dictionary.
		d[word] = definition
	case nil:
		// Word already exists, return an error.
		return ErrWordExists
	default:
		// Return any other unexpected error.
		return err
	}

	// Return nil to indicate success.
	return nil
}

// Update modifies the definition of an existing word in the dictionary.
// The function signature `Update(word, definition string) error` means:
// - It takes two inputs of type `string` (the word and its new definition).
// - It returns a single `error` value, which is nil if the operation is successful.
// This pattern is used to allow the caller to check if the update succeeded or failed.
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		// Word does not exist, so return an error.
		return ErrWordDoesNotExist
	case nil:
		// Word exists, so update the definition.
		d[word] = definition
	default:
		// Return any other unexpected error.
		return err
	}

	// Return nil to indicate success.
	return nil
}

// Delete removes a word from the dictionary.
// The function signature `Delete(word string)` means:
// - It takes a single input of type `string` (the word to delete).
// - It does not return any values, meaning the function cannot fail in a way that needs reporting.
// This is because deleting a non-existent word is harmless, so no error needs to be returned.
func (d Dictionary) Delete(word string) {
	// The built-in delete function removes the word from the map.
	delete(d, word)
}
