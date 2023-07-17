package token

type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF
	// Identifiers (programmar variables)
	// add, foobar, x, y, etc.
	IDENT
	// Integer Literal
	INT
	// = Assignment Operator
	ASSIGN
	// + Addition Operator
	PLUS
	// , Comma Delimiter
	COMMA
	// ; Semicolon Delimiter
	SEMICOLON
	// ( Left Parenthesis
	LPAREN
	// ) Right Parenthesis
	RPAREN
	// { Left Brace
	LBRACE
	// } Right Brace
	RBRACE
	// fn Function keyword
	FUNCTION
	// Let keyword
	LET
)

// String - Creating common behavior - give the type a String function
func (t TokenType) String() string {
	return [...]string{
		"ILLEGAL",
		"EOF",
		"IDENT",
		"INT",
		"ASSIGN",
		"PLUS",
		"COMMA",
		"SEMICOLON",
		"LPAREN",
		"RPAREN",
		"LBRACE",
		"RBRACE",
		"FUNCTION",
		"LET",
	}[t]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (t TokenType) EnumIndex() int {
	return int(t)
}

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
