package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"
	hindi   = "Hindi"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchhHelloPrefix = "Bonjour, "
	hindiHelloPrefix   = "Namaste, "
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchhHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case hindi:
		prefix = hindiHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string {
	if name == "" {
		return englishHelloPrefix + "World"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
