package main

import "fmt"

// Constant for salute
const englishHelloPrefix = "Hello "

// Hello return salute with specific name
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
