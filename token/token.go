package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers + literals
	IDENT = "IDENT" // foobar,x,y,...
	INT   = "INT"   //123455

	//operator
	ASSIGN = "="
	PLUS   = "+"

	//delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	BANG     = "!"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"
	LT       = "<"
	GT       = ">"

	TRUE  = "TRUE"
	FALSE = "FALSE"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	EQ     = "=="
	NOT_EQ = "!="
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
