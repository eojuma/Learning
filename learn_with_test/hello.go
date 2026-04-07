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
	a := [5]int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5, 6}
	c := []int{1, 2, 3, 4, 5}

	fmt.Println(Arrays(a))
	fmt.Println(Slices(c,b))
}

func Hello(name string, language string) string {
	if name == "" {
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

func Repeat() string {
	var res strings.Builder

	for i := 0; i < 5; i++ {
		res.WriteString("A") //This is another way of concatinating strings without 
		// reallocating memory every time a character is added to the string.
	}
	return res.String()
}

func Arrays(a [5]int) int {
	sum := 0

	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func Slices(a,b []int)[]int {
	sum,add:=0,0
	for _, j := range a {
		sum+= j
	}

	for _, j := range b {
		add+=j
	}
	return []int{sum,add}
}
