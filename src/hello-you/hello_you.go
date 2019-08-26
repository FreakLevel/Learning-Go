package main

import "fmt"

// Hello return salute with specific name
func Hello(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(Hello("world"))
}
