package main

import (
	"fmt"
	"os"
)

const (
	LEFT_PARENT  = '('
	RIGHT_PARENT = ')'
	LEFT_BRACE   = '{'
	RIGHT_BRACE  = '}'
	STAR         = '*'
	DOT          = '.'
	COMMA        = ','
	PLUS         = '+'
	MINUS        = '-'
	SEMICOLON    = ';'
	SLASH        = '/'
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	hasError := false

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	for _, char := range fileContents {
		switch char {
		case LEFT_PARENT:
			fmt.Println("LEFT_PAREN ( null")
		case RIGHT_PARENT:
			fmt.Println("RIGHT_PAREN ) null")
		case LEFT_BRACE:
			fmt.Println("LEFT_BRACE { null")
		case RIGHT_BRACE:
			fmt.Println("RIGHT_BRACE } null")
		case STAR:
			fmt.Println("STAR * null")
		case DOT:
			fmt.Println("DOT . null")
		case COMMA:
			fmt.Println("COMMA , null")
		case PLUS:
			fmt.Println("PLUS + null")
		case MINUS:
			fmt.Println("MINUS - null")
		case SEMICOLON:
			fmt.Println("SEMICOLON ; null")
		case SLASH:
			fmt.Println("SLASH / null")
		default:
			hasError = true
			fmt.Fprintf(os.Stderr, "[line 1] Error: Unexpected character: %s", string(char))
		}
	}

	fmt.Println("EOF  null")

	if hasError {
		os.Exit(65)
	}
}
