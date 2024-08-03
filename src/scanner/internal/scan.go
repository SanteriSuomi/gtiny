package internal

import (
	"github.com/SanteriSuomi/gtiny/src/report"
	. "github.com/SanteriSuomi/gtiny/src/token"
	. "github.com/SanteriSuomi/gtiny/src/utils"
)

type Scanner struct {
	src      string
	line     int
	start    int
	curr     int
	tokens   []Token
	keywords map[string]string
	rep      report.ErrorReporter
}

const NEW_LINE = "\n"
const CARRIAGE_RETURN = "\r"
const TAB = "\t"
const DOUBLE_QUOTE = "\""
const WHITE_SPACE = " "
const FRACTION = "."
const EMPTY = ""

func (s *Scanner) Run(input string, rep report.ErrorReporter) []Token {
	s.src = input
	s.rep = rep
	s.tokens = []Token{}
	s.keywords = CreateKeywordMap()
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
	case WHITE_SPACE, CARRIAGE_RETURN, TAB:
		return
	case NEW_LINE:
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
		addToken(IfValueElse(s.nextMatch(EQUAL), BANG_EQUAL, BANG))
	case EQUAL:
		addToken(IfValueElse(s.nextMatch(EQUAL), EQUAL_EQUAL, EQUAL))
	case GREATER:
		addToken(IfValueElse(s.nextMatch(EQUAL), GREATER_EQUAL, GREATER))
	case LESS:
		addToken(IfValueElse(s.nextMatch(EQUAL), LESS_EQUAL, LESS))
	case SLASH:
		// Skip comments
		if s.nextMatch(SLASH) {
			for !s.isAtEnd() && s.peekCurrent() != NEW_LINE {
				s.advance()
			}
		} else {
			addToken(SLASH)
		}
	case STRING:
		s.scanLiteral()
	default:
		if IsDigit(char) {
			s.scanNumber()
		} else if IsLetter(char) {
			s.scanIdentifier()
		} else {
			s.rep.Error(s.line, s.curr, s.curr-s.start, "unexpected character: %s", char)
		}
	}
}

func (s *Scanner) scanLiteral() {
	for (s.peekCurrent() != DOUBLE_QUOTE) && !s.isAtEnd() {
		if s.peekCurrent() == NEW_LINE {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		s.rep.Error(s.line, s.curr, s.curr-s.start, "unterminated string")
	}
	s.advance()

	// Trim the surrounding quotes
	literal := s.src[s.start+1 : s.curr-1]
	s.addTokenLiteral(STRING, literal)
}

func (s *Scanner) scanNumber() {
	for IsDigit(s.peekCurrent()) {
		s.advance()
	}
	if s.peekCurrent() == FRACTION && IsDigit(s.peekNext()) {
		s.advance()
		for IsDigit(s.peekCurrent()) {
			s.advance()
		}
	}
	s.addTokenLiteral(NUMBER, s.src[s.start:s.curr])
}

func (s *Scanner) scanIdentifier() {
	for IsAlphaNumeric(s.peekCurrent()) {
		s.advance()
	}
	literal := s.src[s.start:s.curr]
	keyword := s.keywords[literal]
	if keyword != EMPTY {
		s.addToken(keyword)
	} else {
		s.addTokenLiteral(IDENTIFIER, literal)
	}
}

func (s *Scanner) advance() string {
	char := string(s.src[s.curr])
	s.curr++
	return char
}

func (s *Scanner) peekCurrent() string {
	if s.isAtEnd() {
		return EOF
	}
	return string(s.src[s.curr])
}

func (s *Scanner) peekNext() string {
	if s.nextIsAtEnd() {
		return EOF
	}
	return string(s.src[s.curr+1])
}

func (s *Scanner) nextMatch(expected string) bool {
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

func (s *Scanner) nextIsAtEnd() bool {
	return s.curr+1 >= len(s.src)
}

func (s *Scanner) addToken(tokenType string) {
	s.addTokenLiteral(tokenType, "")
}

func (s *Scanner) addTokenLiteral(tokenType string, literal string) {
	lexeme := s.src[s.start:s.curr]
	length := s.curr - s.start
	s.tokens = append(s.tokens, Token{Type: tokenType, Lexeme: lexeme, Literal: literal, Line: s.line, Column: s.start, Length: length})
}
