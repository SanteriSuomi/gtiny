package entities

type Token struct {
	Type   string
	Value  string
	Line   int
	Column int
	Length int
}

// Special tokens
const (
	EOF = "EOF"
)

// Single-character tokens
const (
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

// Single or double character tokens
const (
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

// Literals
const (
	STRING = "\""
	NUMBER = "NUMBER"
)

// Keywords
const (
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

// CreateKeywordMap returns a map of keywords in the programming language.
// Essentially a set to quickly check if a keyword exists.
//
// Parameters: None
// Returns: A map of strings where the keys and values are the same keywords.
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
