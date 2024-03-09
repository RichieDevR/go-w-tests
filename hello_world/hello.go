package main

import "fmt"

const (
	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishPrefix

	switch language {
	case "Spanish":
		prefix = spanishPrefix
	case "French":
		prefix = frenchPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("", ""))
	fmt.Println(Hello("Elodie", "Spanish"))
	fmt.Println(Hello("Jamie", "French"))
	fmt.Println(Hello("Rich", "English"))
	fmt.Println(Hello("", "Default"))
}
