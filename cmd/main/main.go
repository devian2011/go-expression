package main

import (
	"fmt"
	"log"

	"expressionlng/pkg/token"
)

func main() {
	stack, err := token.ParseFile("./tests/first.lng")
	if err != nil {
		log.Fatalln(err)
	}

	for i, t := range stack.Tokens {
		fmt.Printf("index: %d, value: '%s' Text: '%s'\n", i, string(t.GetVal()), string(t.GetText()))
	}
	fmt.Println()
	fmt.Printf("index: %d", '/')
	fmt.Printf("index: %d", '*')
	fmt.Printf("index: %d", '+')
	fmt.Printf("index: %d", '-')
}
