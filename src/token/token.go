package token

type Token struct {
	Type    string
	Lexeme  string
	Literal string
	Line    int
	Column  int
	Length  int
}

const (
	EOF = "EOF"
)

const (
	// Single-character tokens. These tokens
	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE  = "{"
	RIGHT_BRACE = "}"
	COMMA       = ","
	DOT         = "."
	MINUS       = "-"
	PLUS        = "+"
	SEMICOLON   = ";"
	STAR        = "*"
)

const (
	// Single or double character tokens. Handled after single character tokens by the scanner
	BANG          = "!"
	GREATER       = ">"
	LESS          = "<"
	SLASH         = "/"
	EQUAL         = "="
	BANG_EQUAL    = "!="
	EQUAL_EQUAL   = "=="
	GREATER_EQUAL = ">="
	LESS_EQUAL    = "<="
)

const (
	// Literals
	STRING = "\""
	NUMBER = "NUMBER"
)

const (
	// Identifiers
	IDENTIFIER = "IDENTIFIER"
)

const (
	// Keywords
	AND    = "and"
	CLASS  = "class"
	ELSE   = "else"
	FALSE  = "false"
	FUN    = "fun"
	FOR    = "for"
	IF     = "if"
	NIL    = "nil"
	OR     = "or"
	PRINT  = "print"
	RETURN = "return"
	SUPER  = "super"
	THIS   = "this"
	TRUE   = "true"
	VAR    = "var"
	WHILE  = "while"
)

func CreateKeywordMap() map[string]string {
	return map[string]string{
		AND:    AND,
		CLASS:  CLASS,
		ELSE:   ELSE,
		FALSE:  FALSE,
		FUN:    FUN,
		FOR:    FOR,
		IF:     IF,
		NIL:    NIL,
		OR:     OR,
		PRINT:  PRINT,
		RETURN: RETURN,
		SUPER:  SUPER,
		THIS:   THIS,
		TRUE:   TRUE,
		VAR:    VAR,
		WHILE:  WHILE,
	}
}
