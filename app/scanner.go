package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return '\000'
	}

	return s.source[s.current]
}

func (s *Scanner) string() error {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.", s.line)
		return errors.New("unterminated string")
	}

	s.advance()

	value := string(s.source[s.start+1 : s.current-1])
	s.addTokenLiteral(STRING, value)
	return nil
}

func (s *Scanner) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return '\000'
	}

	return s.source[s.current+1]
}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.advance()

		for s.isDigit(s.peek()) {
			s.advance()
		}
	}

	num, _ := strconv.ParseFloat(string(s.source[s.start:s.current]), 64)

	var float string

	if num == float64(int(num)) {
		float = fmt.Sprintf("%.1f", num)
	} else {
		float = fmt.Sprintf("%g", num)
	}

	s.addTokenLiteral(NUMBER, float)
}

func (s *Scanner) scanTokens() ([]Token, error) {
	hasError := false
	for !s.isAtEnd() {
		s.start = s.current
		err := s.scanToken()

		if err != nil {
			hasError = true
		}
	}

	s.tokens = append(s.tokens, Token{Type: EOF, Lexeme: "", Literal: null, Line: s.line})

	if hasError {
		return s.tokens, errors.New("scanning with errors")
	}

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
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL)
		} else {
			s.addToken(LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL)
		} else {
			s.addToken(GREATER)
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH)
		}
	case ' ', '\t', '\r':
	case '\n':
		s.line++
	case '"':
		if err := s.string(); err != nil {
			return err
		}
	default:
		if s.isDigit(c) {
			s.number()
		} else {
			fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %s\n", s.line, string(c))
			return errors.New("lexical error")
		}
	}

	return nil
}
