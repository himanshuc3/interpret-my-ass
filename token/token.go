package token

// NOTE:
// 1. const ILLEGAL string = "ILLEGAL" - explicit const
// Btw, very og syntax with let and const. Well, ackshually
// technically, it was included in ES6 spec in JS, so it might
// have been copied from golang to js
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// NOTE:
// 1. In go, a const can only be a compile time constant like
// number string, boolean
var TokenMapping map[TokenType]Token = GenerateTokenMapping()

// NOTE:
// 1. token type is declared as string instead of int/byte which
// might be better for performance and provide space gains
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// NOTE:
// 1. Can be removed in favor of a global variable since global
// map constants aren't allowed in go
func GenerateTokenMapping() map[TokenType]Token {

	// NOTE:
	// 1. Dynamic property names like in JS [COMMA]
	tokens := map[TokenType]Token{
		ASSIGN: {
			Type: ASSIGN,
		},
		SEMICOLON: {
			Type: SEMICOLON,
		},
		LPAREN: {
			Type: LPAREN,
		},
		COMMA: {
			Type: COMMA,
		},
		PLUS: {
			Type: PLUS,
		},
		LBRACE: {
			Type: LBRACE,
		},
		RBRACE: {
			Type: RBRACE,
		},
	}
	return tokens
}
