package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"

	englishHello = "Hello "
	spanishHello = "Hola "
	frenchHello  = "Bonjour "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World!"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHello
	case french:
		prefix = frenchHello
	default:
		prefix = englishHello
	}
	return
}

func main() {
	fmt.Println(Hello("Dashiell!", "English"))
}
