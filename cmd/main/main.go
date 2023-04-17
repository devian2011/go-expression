package main

import (
	"fmt"
	"log"

	"expressionlng/src/token"
)

func main() {
	stack, err := token.ParseFile("./tests/first.lng")
	if err != nil {
		log.Fatalln(err)
	}

	for i, t := range stack.Tokens {
		fmt.Printf("index: %d, type: %d value: '%s' Text: '%s'\n", i, t.GetType(), string(t.GetVal()))
	}
	fmt.Println()
	fmt.Printf("index: %d", '/')
	fmt.Printf("index: %d", '*')
	fmt.Printf("index: %d", '+')
	fmt.Printf("index: %d", '-')
}
