package main

import "fmt"

const (
	portuguese = "Portuguese"
	french     = "French"

	englishHelloPrefix    = "Hello, "
	portugueseHelloPrefix = "Olá, "
	frenchHelloPrefix     = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case portuguese:
		prefix = portugueseHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", "English"))
}
