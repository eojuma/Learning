package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		 greeter("Hello")
		 greeter("World")
	}
}

func greeter(s string) {
	fmt.Println(s)
}
