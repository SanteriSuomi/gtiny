package internal

import (
	"fmt"

	. "github.com/SanteriSuomi/gtiny/src/error"
	. "github.com/SanteriSuomi/gtiny/src/token"
)

type Scanner struct {
	source   string
	line     int
	start    int
	current  int
	tokens   []Token
	reporter ErrorReporter
}

func (s *Scanner) Run(input string, reporter ErrorReporter) {
	s.source = input
	s.reporter = reporter
	s.scanTokens()
	for _, token := range s.tokens {
		fmt.Println(token.Type)
	}
}

func (s *Scanner) scanTokens() {
	s.tokens = []Token{}
	s.start, s.current, s.line = 0, 0, 1
	s.start = s.current
	for s.current < len(s.source) {
		s.scanToken()
	}
}

func (s *Scanner) scanToken() {
	c := string(s.source[s.current])
	s.current++
	// Single character tokens
	switch c {
	case LEFT_PAREN:
		s.addToken(LEFT_PAREN)
	case RIGHT_PAREN:
		s.addToken(RIGHT_PAREN)
	case LEFT_BRACE:
		s.addToken(LEFT_BRACE)
	case RIGHT_BRACE:
		s.addToken(RIGHT_BRACE)
	case COMMA:
		s.addToken(COMMA)
	case DOT:
		s.addToken(DOT)
	case MINUS:
		s.addToken(MINUS)
	case PLUS:
		s.addToken(PLUS)
	case SEMICOLON:
		s.addToken(SEMICOLON)
	case SLASH:
		s.addToken(SLASH)
	case STAR:
		s.addToken(STAR)
	case BANG:
		s.addToken(BANG)
	default:
		s.reporter.Error(s.line, s.current, s.current-s.start, fmt.Sprintf("Unexpected character: %s", c))
	}
}

func (s *Scanner) addToken(tokenType string) {
	lexeme := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{Type: tokenType, Lexeme: lexeme, Literal: "", Line: s.line, Column: s.start, Length: s.current - s.start})
}
