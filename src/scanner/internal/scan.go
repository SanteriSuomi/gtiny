package internal

import (
	"github.com/SanteriSuomi/gtiny/src/report"
	. "github.com/SanteriSuomi/gtiny/src/token"
	. "github.com/SanteriSuomi/gtiny/src/utils"
)

type Scanner struct {
	src    string
	line   int
	start  int
	curr   int
	tokens []Token
	rep    report.ErrorReporter
}

func (s *Scanner) Run(input string, rep report.ErrorReporter) []Token {
	s.src = input
	s.rep = rep
	s.tokens = []Token{}
	s.scanTokens()
	s.addToken(EOF)
	return s.tokens
}

func (s *Scanner) scanTokens() {
	s.start, s.curr, s.line = 0, 0, 0
	for !s.isAtEnd() {
		s.scanToken()
		s.start = s.curr
	}
}

func (s *Scanner) scanToken() {
	addToken := s.addToken
	char := s.advance()

	switch char {
	case " ", "\r", "\t":
		return
	case "\n":
		s.line++
		return
	// Single character tokens
	case LEFT_PAREN:
		addToken(LEFT_PAREN)
	case RIGHT_PAREN:
		addToken(RIGHT_PAREN)
	case LEFT_BRACE:
		addToken(LEFT_BRACE)
	case RIGHT_BRACE:
		addToken(RIGHT_BRACE)
	case COMMA:
		addToken(COMMA)
	case DOT:
		addToken(DOT)
	case MINUS:
		addToken(MINUS)
	case PLUS:
		addToken(PLUS)
	case SEMICOLON:
		addToken(SEMICOLON)
	case STAR:
		addToken(STAR)

	// Single or double character tokes
	case BANG:
		addToken(IfElse(s.match(EQUAL), BANG_EQUAL, BANG))
	case EQUAL:
		addToken(IfElse(s.match(EQUAL), EQUAL_EQUAL, EQUAL))
	case GREATER:
		addToken(IfElse(s.match(EQUAL), GREATER_EQUAL, GREATER))
	case LESS:
		addToken(IfElse(s.match(EQUAL), LESS_EQUAL, LESS))
	case SLASH:
		// Skip comments
		if s.match(SLASH) {
			for !s.isAtEnd() && s.peek() != "\n" {
				s.advance()
			}
		} else {
			addToken(SLASH)
		}

	// TODO: Literals
	case STRING:
		s.addTokenLiteral(STRING, "")

	default:
		s.rep.Error(s.line, s.curr, s.curr-s.start, "unexpected character: %s", char)
	}
}

func (s *Scanner) advance() string {
	char := string(s.src[s.curr])
	s.curr++
	return char
}

func (s *Scanner) peek() string {
	return string(s.src[s.curr])
}

func (s *Scanner) match(expected string) bool {
	if s.curr >= len(s.src) {
		return false
	}
	if string(s.src[s.curr]) == expected {
		s.curr++
		return true
	}
	return false
}

func (s *Scanner) isAtEnd() bool {
	return s.curr >= len(s.src)
}

func (s *Scanner) addToken(tokenType string) {
	s.addTokenLiteral(tokenType, "")
}

func (s *Scanner) addTokenLiteral(tokenType string, literal string) {
	lexeme := s.src[s.start:s.curr]
	length := s.curr - s.start
	s.tokens = append(s.tokens, Token{Type: tokenType, Lexeme: lexeme, Literal: literal, Line: s.line, Column: s.start, Length: length})
}
