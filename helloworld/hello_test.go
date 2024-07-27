package main

import "testing"

// TestHello contains subtests to verify the Hello function
func TestHello(t *testing.T) {
	// Subtest: Saying hello to people
	t.Run("saying hello to people", func(t *testing.T) {
		// Call the Hello function
		got := Hello("Aniello", "")
		// Expected result
		want := "Hello, Aniello"
		// Verify the result
		assertCorrectMessage(t, got, want)
	})

	// Subtest: Default to "World" if the name is an empty string
	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	// Subtest: Greeting in Spanish
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	//Subtest: French
	t.Run("in French", func(t *testing.T) {
		got := Hello("Aniello", "French")
		want := "Bonjour, Aniello"
		assertCorrectMessage(t, got, want)
	})
}

// Helper function to verify if the got and want strings are equal
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
