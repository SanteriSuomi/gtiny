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
	EOF string = "EOF"
)

const (
	// Single-character tokens
	LEFT_PAREN  string = "("
	RIGHT_PAREN string = ")"
	LEFT_BRACE  string = "{"
	RIGHT_BRACE string = "}"
	COMMA       string = ","
	DOT         string = "."
	MINUS       string = "-"
	PLUS        string = "+"
	SEMICOLON   string = ";"
	SLASH       string = "/"
	STAR        string = "*"
)

const (
	// One or two character tokens
	BANG          string = "!"
	BANG_EQUAL    string = "!="
	EQUAL         string = "="
	EQUAL_EQUAL   string = "=="
	GREATER       string = ">"
	GREATER_EQUAL string = ">="
	LESS          string = "<"
	LESS_EQUAL    string = "<="
)

const (
	// Literals
	IDENTIFIER string = "IDENTIFIER"
	STRING     string = "STRING"
	NUMBER     string = "NUMBER"
)

const (
	// Keywords
	AND    string = "and"
	CLASS  string = "class"
	ELSE   string = "else"
	FALSE  string = "false"
	FUN    string = "fun"
	FOR    string = "fo"
	IF     string = "if"
	NIL    string = "nil"
	OR     string = "or"
	PRINT  string = "print"
	RETURN string = "return"
	SUPER  string = "super"
	THIS   string = "this"
	TRUE   string = "true"
	VAR    string = "var"
	WHILE  string = "while"
)
