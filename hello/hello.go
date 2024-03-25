package main

import "fmt"

const englishHelloPrefix = "Hello, "
const portugueseHelloPrefix = "Olá, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := englishHelloPrefix

	switch language {
	case "Portuguese":
		prefix = portugueseHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	}

	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("world", "English"))
}
