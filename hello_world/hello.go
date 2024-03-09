package main

import "fmt"

const (
	englishPrefix    = "Hello, "
	spanishPrefix    = "Hola, "
	frenchPrefix     = "Bonjour, "
	portuguesePrefix = "Olá, "
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
	case "Portuguese":
		prefix = portuguesePrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("", ""))
	fmt.Println(Hello("Ricardo", "Spanish"))
	fmt.Println(Hello("Ricardo", "Portuguese"))
	fmt.Println(Hello("Richard", "French"))
	fmt.Println(Hello("Richie", "English"))
	fmt.Println(Hello("", "Default"))
}
