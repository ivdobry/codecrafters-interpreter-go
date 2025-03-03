package main

import (
	"errors"
	"fmt"
	"os"
)

type TokenType int

const (
	EOF TokenType = iota
	LEFT_PARENT
	RIGHT_PARENT
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	IDENTIFIER
	STRING
	NUMBER
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	null
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

func (t TokenType) String() string {
	return [...]string{
		"EOF",
		"LEFT_PAREN", "RIGHT_PAREN", "LEFT_BRACE", "RIGHT_BRACE",
		"COMMA", "DOT", "MINUS", "PLUS",
		"SEMICOLON", "SLASH", "STAR",
		"BANG", "BANG_EQUAL",
		"EQUAL", "EQUAL_EQUAL",
		"GREATER", "GREATER_EQUAL",
		"LESS", "LESS_EQUAL",
		"IDENTIFIER", "STRING",
		"NUMBER",
		"AND", "CLASS", "ELSE", "FALSE", "FUN", "FOR", "IF", "null", "OR",
		"PRINT", "RETURN", "SUPER", "THIS", "TRUE", "VAR", "WHILE",
	}[t]
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %s %s", t.Type, t.Lexeme, t.Literal)
}

type Scanner struct {
	source  []byte
	tokens  []Token
	start   int
	current int
	line    int
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) addToken(token TokenType) {
	s.addTokenLiteral(token, null)
}

func (s *Scanner) addTokenLiteral(token TokenType, literal interface{}) {
	text := string(s.source[s.start:s.current])
	s.tokens = append(s.tokens, Token{Type: token, Lexeme: text, Literal: literal, Line: s.line})
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) advance() byte {
	s.current++
	return s.source[s.current-1]
}

func (s *Scanner) scanTokens() ([]Token, error) {
	for !s.isAtEnd() {
		s.start = s.current
		err := s.scanToken()

		if err != nil {
			return s.tokens, errors.New("scanning with errors")
		}
	}

	s.tokens = append(s.tokens, Token{Type: EOF, Lexeme: "", Literal: null, Line: s.line})
	return s.tokens, nil
}

func (s *Scanner) scanToken() error {
	c := s.advance()

	switch c {
	case '(':
		s.addToken(LEFT_PARENT)
	case ')':
		s.addToken(RIGHT_PARENT)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}
	default:
		fmt.Fprintf(os.Stderr, "[line 1] Error: Unexpected character: %s\n", string(c))
		return errors.New("lexical error")
	}

	return nil
}

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

	for _, token := range tokens {
		result = result + token.String()
	}

	fmt.Println(result)

	if err != nil {
		os.Exit(65)
	}
}
