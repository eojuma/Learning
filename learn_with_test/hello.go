package main

import (
	"fmt"
	"strings"
)

const englishPrefix = "Hello, "
const spanishprefix = "Hola, "
const spanish = "spanish,"
const french = "french"
const frenchprefix = "Bonjour,"
const english = "english"

func main() {
	fmt.Println(Hello("", ""))
	fmt.Println(Add(4, 5))
	// Expected  output:9
	fmt.Println(Repeat())
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

func Add(a, b int) int {
	return a + b
}

func Repeat()string{
	var res strings.Builder

	for i:=0;i<5;i++{
		res.WriteString("A")
	}
	return res.String()
}
