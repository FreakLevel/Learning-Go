package main

import "fmt"

// Constant for salute
const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "

// Hello return salute with specific name
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return languagePrefix(language) + name
}

func languagePrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
