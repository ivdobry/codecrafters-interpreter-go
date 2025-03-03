package main

import (
	"fmt"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

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

	scanner := &Scanner{
		source:  fileContents,
		start:   0,
		current: 0,
		line:    1,
	}

	tokens, err := scanner.scanTokens()

	var result string

	for i, token := range tokens {
		if i == 0 {
			result = token.String()
		} else {
			result = result + "\n" + token.String()
		}
	}

	fmt.Println(result)

	if err != nil {
		os.Exit(65)
	}
}
