package main

import "fmt"

const englishHelloPrefix = "Hello, "
const russianHelloPrefix = "Привет, "
const spanishHelloPrefix = "Hola, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return getLanguagePrefix(language) + name
}

func getLanguagePrefix(language string) string {
	switch language {
	case "english":
		return englishHelloPrefix
	case "spanish":
		return spanishHelloPrefix
	case "russian":
		return russianHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func main() {
	fmt.Println(Hello("Diego", "english"))
}
