package main

import (
	"fmt"
)

const englishPrefix = "Hello, "
const spanishprefix = "Hola, "
const spanish = "spanish,"
const french = "french"
const frenchprefix = "Bonjour,"
const english = "english"

func main() {
	fmt.Println(Hello("", ""))
}

func Hello(name string, language string) string {
	if name == "" {
		language = "english"
		name = "World!"
	}
	if language == "spanish" {
		return spanishprefix + name
	} else if language == "french" {
		return frenchprefix + name
	} else {

		return englishPrefix + name
	}
}
