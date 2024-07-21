package token

const (
	// Single-character tokens
	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE  = "{"
	RIGHT_BRACE = "}"
	COMMA       = ","
	DOT         = "."
	MINUS       = "-"
	PLUS        = "+"
	SEMICOLON   = ";"
	SLASH       = "/"
	STAR        = "*"
)

const (
	// One or two character tokens
	BANG          = "!"
	BANG_EQUAL    = "!="
	EQUAL         = "="
	EQUAL_EQUAL   = "=="
	GREATER       = ">"
	GREATER_EQUAL = ">="
	LESS          = "<"
	LESS_EQUAL    = "<="
)

const (
	// Literals
	IDENTIFIER = "IDENTIFIER"
	STRING     = "STRING"
	NUMBER     = "NUMBER"
)

const (
	// Keywords
	AND    = "and"
	CLASS  = "class"
	ELSE   = "else"
	FALSE  = "false"
	FUN    = "fun"
	FOR    = "fo"
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

type Token struct {
	Type    string
	Lexeme  string
	Literal string
	Line    int
}
