package token

// TokenType holds the lexical type information
type TokenType string

// Token holds the lexical token information
type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"else":   ELSE,
	"false":  FALSE,
	"fn":     FUNCTION,
	"if":     IF,
	"let":    LET,
	"return": RETURN,
	"true":   TRUE,
}

// LookupIdent check if the ident is part of the
// language keyword list
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN   = "="
	EQ       = "=="
	NOT_EQ   = "=="
	ASTERISK = "*"
	BANG     = "!"
	MINUS    = "-"
	PLUS     = "+"
	SLASH    = "/"

	GT = ">"
	LT = "<"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"

	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
)
