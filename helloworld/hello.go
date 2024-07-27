package main

import "fmt"

// Constants to define supported languages and their greeting prefixes
const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// Hello function returns a greeting for the given name and language
func Hello(name string, language string) string {
	// Default name to "World" if no name is provided
	if name == "" {
		name = "World"
	}

	// Use the greetingPrefix function to get the appropriate greeting prefix
	return greetingPrefix(language) + name
}

// greetingPrefix function returns the appropriate greeting prefix based on the language
func greetingPrefix(language string) (prefix string) {
	// Use a switch statement to determine the greeting prefix
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	// Return the determined prefix
	return
}

func main() {
	// Call the Hello function with a name and language and print the result
	fmt.Println(Hello("Aniello", "French"))
}
